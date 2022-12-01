package tables

import (
	"context"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudSecurityCenterVersionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudSecurityCenterVersionGenerator{}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetTableName() string {
	return "alicloud_security_center_version"
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*alicloud_client.AliCloudClient).Region

			supportedRegions := map[string]struct{}{"cn-hangzhou": {}, "ap-southeast-1": {}, "ap-southeast-3": {}}
			if _, exists := supportedRegions[region]; !exists {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			client, err := alicloud_client.SecurityCenterService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := sas.CreateDescribeVersionConfigRequest()
			request.Scheme = "https"

			response, err := client.DescribeVersionConfig(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- versionInfo{*response, region}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type versionInfo struct {
	sas.DescribeVersionConfigResponse
	Region string
}

func getSecurityCenterVersionAkas(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(versionInfo)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:security-center:" + data.Region + ":" + accountID + ":version/" + strconv.Itoa(data.Version)}

	return akas, nil
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("asset_level").ColumnType(schema.ColumnTypeInt).Description("The purchased quota for Security Center.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_over_balance").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the number of existing servers exceeds your quota.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_defined_alarms").ColumnType(schema.ColumnTypeInt).Description("Indicates whether the custom alert feature is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the purchased Security Center instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_trial_version").ColumnType(schema.ColumnTypeBool).Description("Indicates whether Security Center is the free trial edition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_white_list_auth_count").ColumnType(schema.ColumnTypeInt).Description("The quota on the servers to which you can apply your application whitelist.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_trail_end_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the last free trial ends.").
			Extractor(column_value_extractor.StructSelector("LastTrailEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("release_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the Security Center instance expired.").
			Extractor(column_value_extractor.StructSelector("ReleaseTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sas_log").ColumnType(schema.ColumnTypeInt).Description("Indicates whether log analysis is purchased.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getSecurityCenterVersionAkas(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Description("The purchased edition of Security Center.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_lock_auth_count").ColumnType(schema.ColumnTypeInt).Description("The quota on the servers that web tamper proofing protects.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sas_screen").ColumnType(schema.ColumnTypeInt).Description("Indicates whether the security dashboard is purchased.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sls_capacity").ColumnType(schema.ColumnTypeInt).Description("The purchased capacity of log storage.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_lock").ColumnType(schema.ColumnTypeInt).Description("Indicates whether web tamper proofing is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_white_list").ColumnType(schema.ColumnTypeInt).Description("Indicates whether the application whitelist is enabled.").Build(),
	}
}

func (x *TableAlicloudSecurityCenterVersionGenerator) GetSubTables() []*schema.Table {
	return nil
}
