package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudEcsAutoProvisioningGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsAutoProvisioningGroupGenerator{}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetTableName() string {
	return "alicloud_ecs_auto_provisioning_group"
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeAutoProvisioningGroupsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeAutoProvisioningGroups(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, group := range response.AutoProvisioningGroups.AutoProvisioningGroup {
					resultChannel <- group
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

func ecsAutosProvisioningGroupTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(ecs.AutoProvisioningGroup)

	title := data.AutoProvisioningGroupId

	if len(data.AutoProvisioningGroupName) > 0 {
		title = data.AutoProvisioningGroupName
	}

	return title, nil
}
func getEcsAutosProvisioningGroupAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ecs.AutoProvisioningGroup)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ecs:" + data.RegionId + ":" + accountID + ":auto-provisioning-group/" + data.AutoProvisioningGroupId}

	return akas, nil
}
func getEcsAutosProvisioningGroupInstances(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ecs.AutoProvisioningGroup)

	client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ecs.CreateDescribeAutoProvisioningGroupInstancesRequest()
	request.Scheme = "https"
	request.AutoProvisioningGroupId = data.AutoProvisioningGroupId

	response, err := client.DescribeAutoProvisioningGroupInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response, nil
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_provisioning_group_type").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocation_strategy").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").
			Extractor(column_value_extractor.StructSelector("PayAsYouGoOptions.AllocationStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("terminate_instances").ColumnType(schema.ColumnTypeBool).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_options").ColumnType(schema.ColumnTypeJSON).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsAutosProvisioningGroupAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A friendly name for the resource.").
			Extractor(column_value_extractor.StructSelector("AutoProvisioningGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_provisioning_group_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_from").ColumnType(schema.ColumnTypeTimestamp).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_until").ColumnType(schema.ColumnTypeTimestamp).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template_configs").ColumnType(schema.ColumnTypeJSON).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template_version").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_spot_price").ColumnType(schema.ColumnTypeFloat).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("terminate_instances_with_expiration").ColumnType(schema.ColumnTypeBool).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("excess_capacity_termination_policy").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instances").ColumnType(schema.ColumnTypeJSON).Description("An unique identifier for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsAutosProvisioningGroupInstances(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Instances.Instance")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_capacity_specification").ColumnType(schema.ColumnTypeJSON).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := ecsAutosProvisioningGroupTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudEcsAutoProvisioningGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
