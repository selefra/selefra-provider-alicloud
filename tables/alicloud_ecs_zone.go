package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudEcsZoneGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsZoneGenerator{}

func (x *TableAlicloudEcsZoneGenerator) GetTableName() string {
	return "alicloud_ecs_zone"
}

func (x *TableAlicloudEcsZoneGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsZoneGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsZoneGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsZoneGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*alicloud_client.AliCloudClient).Region

			client, err := alicloud_client.ECSRegionService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := ecs.CreateDescribeZonesRequest()
			request.Scheme = "https"
			request.RegionId = region
			request.AcceptLanguage = "en-US"

			response, err := client.DescribeZones(request)

			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range response.Zones.Zone {
				resultChannel <- zoneInfo{i, region}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type zoneInfo = struct {
	ecs.Zone
	Region string
}

func getZoneAkas(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(zoneInfo)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ecs::" + accountID + ":zone/" + data.ZoneId}, nil
}

func (x *TableAlicloudEcsZoneGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsZoneGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getZoneAkas(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The zone ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_disk_categories").ColumnType(schema.ColumnTypeJSON).Description("The supported disk categories. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("AvailableDiskCategories.DiskCategories")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_host_generations").ColumnType(schema.ColumnTypeJSON).Description("The generation numbers of dedicated hosts. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("DedicatedHostGenerations.DedicatedHostGeneration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_resource_creation").ColumnType(schema.ColumnTypeJSON).Description("The types of the resources that can be created. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("AvailableResourceCreation.ResourceTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_volume_categories").ColumnType(schema.ColumnTypeJSON).Description("The categories of available shared storage. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("AvailableVolumeCategories.VolumeCategories")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("local_name").ColumnType(schema.ColumnTypeString).Description("The name of the zone in the local language.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_dedicated_host_types").ColumnType(schema.ColumnTypeJSON).Description("The supported types of dedicated hosts. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("AvailableDedicatedHostTypes.DedicatedHostType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_instance_types").ColumnType(schema.ColumnTypeJSON).Description("The instance types of instances that can be created. The data type of this parameter is List.").
			Extractor(column_value_extractor.StructSelector("AvailableInstanceTypes.InstanceTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_resources").ColumnType(schema.ColumnTypeJSON).Description("An array consisting of ResourcesInfo data.").
			Extractor(column_value_extractor.StructSelector("AvailableResources.ResourcesInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alicloud_ecs_region_selefra_id").ColumnType(schema.ColumnTypeString).Description("fk to alicloud_ecs_region.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAlicloudEcsZoneGenerator) GetSubTables() []*schema.Table {
	return nil
}
