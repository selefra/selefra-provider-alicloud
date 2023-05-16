package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudEcsImageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsImageGenerator{}

func (x *TableAlicloudEcsImageGenerator) GetTableName() string {
	return "alicloud_ecs_image"
}

func (x *TableAlicloudEcsImageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsImageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsImageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsImageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeImagesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeImages(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, image := range response.Images.Image {
					resultChannel <- image
					count++
				}
				if count >= response.TotalCount {
					break
				}
				request.PageNumber = requests.NewInteger(response.PageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getEcsImageARN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ecs.Image)
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + region + ":" + accountID + ":image/" + data.ImageId

	return arn, nil
}
func getEcsImageRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	return region, nil
}
func getEcsImageSharePermission(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ecs.Image)

	if data.ImageOwnerAlias != "self" {
		return nil, nil
	}

	id := data.ImageId

	client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}
	request := ecs.CreateDescribeImageSharePermissionRequest()
	request.Scheme = "https"
	request.ImageId = id

	var groups []ecs.ShareGroup
	var accounts []ecs.Account

	count := 0
	for {
		response, err := client.DescribeImageSharePermission(request)
		if err != nil {

			return nil, err
		}
		for _, group := range response.ShareGroups.ShareGroup {

			groups = append(groups, group)
			count++
		}
		for _, account := range response.Accounts.Account {

			accounts = append(accounts, account)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}

	r := map[string]interface{}{
		"ShareGroups": groups,
		"Accounts":    accounts,
	}

	return r, nil
}

func (x *TableAlicloudEcsImageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsImageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("os_name_en").ColumnType(schema.ColumnTypeString).Description("The English name of the operating system.").
			Extractor(column_value_extractor.StructSelector("OSNameEn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("ImageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_family").ColumnType(schema.ColumnTypeString).Description("The name of the image family.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_subscribed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether you have subscribed to the image that corresponds to the specified product code.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).Description("The platform of the operating system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("progress").ColumnType(schema.ColumnTypeString).Description("The image creation progress, in percent(%).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsImageRegion(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeString).Description("The image architecture. Possible values are: 'i386', and 'x86_64'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_copied").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the image is a copy of another image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_self_shared").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the image has been shared to other Alibaba Cloud accounts.").
			Extractor(column_value_extractor.StructSelector("Image.IsSelfShared")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsImageARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage").ColumnType(schema.ColumnTypeString).Description("Indicates whether the image has been used to create ECS instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A friendly name of the resource.").
			Extractor(column_value_extractor.StructSelector("ImageName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_version").ColumnType(schema.ColumnTypeString).Description("The version of the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_support_cloud_init").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the image supports cloud-init.").
			Extractor(column_value_extractor.StructSelector("IsSupportCloudinit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_name").ColumnType(schema.ColumnTypeString).Description("The Chinese name of the operating system.").
			Extractor(column_value_extractor.StructSelector("OSName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_type").ColumnType(schema.ColumnTypeString).Description("The type of the operating system. Possible values are: windows,and linux").
			Extractor(column_value_extractor.StructSelector("OSType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_device_mappings").ColumnType(schema.ColumnTypeJSON).Description("The mappings between disks and snapshots under the image.").
			Extractor(column_value_extractor.StructSelector("DiskDeviceMappings.DiskDeviceMapping")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("share_permissions").ColumnType(schema.ColumnTypeJSON).Description("A list of groups and accounts that the image can be shared.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsImageSharePermission(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).Description("The ID of the image that the instance is running.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the ECS image.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsImageARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_owner_alias").ColumnType(schema.ColumnTypeString).Description("The alias of the image owner. Possible values are: system, self, others, marketplace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the image was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_code").ColumnType(schema.ColumnTypeString).Description("The product code of the Alibaba Cloud Marketplace image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the image.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The size of the image (in GiB).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A user-defined, human readable description for the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_support_io_optimized").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the image can be used on I/O optimized instances.").
			Extractor(column_value_extractor.StructSelector("IsSupportIoOptimized")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the image belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudEcsImageGenerator) GetSubTables() []*schema.Table {
	return nil
}
