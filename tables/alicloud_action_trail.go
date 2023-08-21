package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/actiontrail"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudActionTrailGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudActionTrailGenerator{}

func (x *TableAlicloudActionTrailGenerator) GetTableName() string {
	return "alicloud_action_trail"
}

func (x *TableAlicloudActionTrailGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudActionTrailGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudActionTrailGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudActionTrailGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ActionTrailService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := actiontrail.CreateDescribeTrailsRequest()
			request.Scheme = "https"
			request.IncludeShadowTrails = requests.NewBoolean(true)

			response, err := client.DescribeTrails(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, trail := range response.TrailList {
				resultChannel <- trail
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getActionTrailAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(actiontrail.TrailListItem)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:actiontrail:" + data.HomeRegion + ":" + accountID + ":actiontrail/" + data.Name}

	return akas, nil
}
func getActionTrailRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	return region, nil
}

func (x *TableAlicloudActionTrailGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudActionTrailGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_organization_trail").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the trail was created as a multi-account trail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_rw").ColumnType(schema.ColumnTypeString).Description("The read/write type of the delivered events.").
			Extractor(column_value_extractor.StructSelector("EventRW")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sls_project_arn").ColumnType(schema.ColumnTypeString).Description("The ARN of the Log Service project to which events are delivered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_logging_time").ColumnType(schema.ColumnTypeTimestamp).Description("The most recent date and time when logging was enabled for the trail.").
			Extractor(column_value_extractor.StructSelector("StartLoggingTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the trail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Description("The name of the Resource Access Management (RAM) role that ActionTrail is allowed to assume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the trail was created.").
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sls_write_role_arn").ColumnType(schema.ColumnTypeString).Description("The ARN of the RAM role assumed by ActionTrail for delivering logs to the destination Log Service project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trail_region").ColumnType(schema.ColumnTypeString).Description("The regions to which the trail is applied.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getActionTrailAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("home_region").ColumnType(schema.ColumnTypeString).Description("The home region of the trail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("oss_key_prefix").ColumnType(schema.ColumnTypeString).Description("The prefix of log files stored in the OSS bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stop_logging_time").ColumnType(schema.ColumnTypeTimestamp).Description("The most recent date and time when logging was disabled for the trail.").
			Extractor(column_value_extractor.StructSelector("StopLoggingTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).Description("The most recent time when the configuration of the trail was updated.").
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getActionTrailRegion(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the trail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("oss_bucket_name").ColumnType(schema.ColumnTypeString).Description("The name of the OSS bucket to which events are delivered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAlicloudActionTrailGenerator) GetSubTables() []*schema.Table {
	return nil
}
