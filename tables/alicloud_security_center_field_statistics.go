package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/turbot/go-kit/helpers"
)

type TableAlicloudSecurityCenterFieldStatisticsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudSecurityCenterFieldStatisticsGenerator{}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetTableName() string {
	return "alicloud_security_center_field_statistics"
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*alicloud_client.AliCloudClient).Region

			supportedRegions := []string{"cn-hangzhou", "ap-southeast-1", "ap-southeast-3"}
			if !helpers.StringSliceContains(supportedRegions, region) {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			client, err := alicloud_client.SecurityCenterService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := sas.CreateDescribeFieldStatisticsRequest()
			request.Scheme = "https"

			response, err := client.DescribeFieldStatistics(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- FieldInfo{response.GroupedFields, region}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

var supportedRegions = []string{"cn-hangzhou", "ap-south-1", "me-east-1", "eu-central-1", "ap-northeast-1", "ap-southeast-2"}

type FieldInfo struct {
	sas.GroupedFields
	Region string
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("general_asset_count").ColumnType(schema.ColumnTypeInt).Description("The number of general assets.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("important_asset_count").ColumnType(schema.ColumnTypeInt).Description("The number of important assets.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("offline_instance_count").ColumnType(schema.ColumnTypeInt).Description("The number of offline servers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unprotected_instance_count").ColumnType(schema.ColumnTypeInt).Description("The number of unprotected assets.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("new_instance_count").ColumnType(schema.ColumnTypeInt).Description("The number of new servers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_count").ColumnType(schema.ColumnTypeInt).Description("The number of regions to which the servers belong.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category_count").ColumnType(schema.ColumnTypeInt).Description("The number of assets category.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_count").ColumnType(schema.ColumnTypeInt).Description("The number of asset groups.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_count").ColumnType(schema.ColumnTypeInt).Description("The total number of assets of the specified type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_running_status_count").ColumnType(schema.ColumnTypeInt).Description("The number of inactive servers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("risk_instance_count").ColumnType(schema.ColumnTypeInt).Description("The number of assets that are at risk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("test_asset_count").ColumnType(schema.ColumnTypeInt).Description("The number of test assets.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_count").ColumnType(schema.ColumnTypeInt).Description("The number of VPCs.").Build(),
	}
}

func (x *TableAlicloudSecurityCenterFieldStatisticsGenerator) GetSubTables() []*schema.Table {
	return nil
}
