package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudRamPasswordPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamPasswordPolicyGenerator{}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetTableName() string {
	return "alicloud_ram_password_policy"
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ram.CreateGetPasswordPolicyRequest()
			request.Scheme = "https"
			response, err := client.GetPasswordPolicy(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- response.PasswordPolicy
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("hard_expiry").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the password has expired.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_lowercase_characters").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a password must contain one or more lowercase letters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_uppercase_characters").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a password must contain one or more uppercase letters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_login_attempts").ColumnType(schema.ColumnTypeInt).Description("The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.").
			Extractor(column_value_extractor.StructSelector("MaxLoginAttemps")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_password_age").ColumnType(schema.ColumnTypeInt).Description("The number of days for which a password is valid. Default value: 0. The default value indicates that the password never expires.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_password_length").ColumnType(schema.ColumnTypeInt).Description("The minimum required number of characters in a password.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_reuse_prevention").ColumnType(schema.ColumnTypeInt).Description("The number of previous passwords that the user is prevented from reusing. Default value: 0. The default value indicates that the RAM user is not prevented from reusing previous passwords.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_numbers").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a password must contain one or more digits.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_symbols").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a password must contain one or more special characters.").Build(),
	}
}

func (x *TableAlicloudRamPasswordPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
