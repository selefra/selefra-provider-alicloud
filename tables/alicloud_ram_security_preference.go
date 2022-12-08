package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudRamSecurityPreferenceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamSecurityPreferenceGenerator{}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetTableName() string {
	return "alicloud_ram_security_preference"
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ram.CreateGetSecurityPreferenceRequest()
			request.Scheme = "https"
			response, err := client.GetSecurityPreference(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- response.SecurityPreference
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("enable_save_mfa_ticket").ColumnType(schema.ColumnTypeBool).Description("Indicates whether RAM users can save security codes for multi-factor authentication (MFA) during logon. Each security code is valid for seven days.").
			Extractor(column_value_extractor.StructSelector("LoginProfilePreference.EnableSaveMFATicket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login_session_duration").ColumnType(schema.ColumnTypeInt).Description("The validity period of a logon session of a RAM user. Unit: hours.").
			Extractor(column_value_extractor.StructSelector("LoginProfilePreference.LoginSessionDuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_user_to_change_password").ColumnType(schema.ColumnTypeBool).Description("Indicates whether RAM users can change their passwords.").
			Extractor(column_value_extractor.StructSelector("LoginProfilePreference.AllowUserToChangePassword")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_user_to_manage_public_keys").ColumnType(schema.ColumnTypeBool).Description("Indicates whether RAM users can manage their public keys.").
			Extractor(column_value_extractor.StructSelector("PublicKeyPreference.AllowUserToManagePublicKeys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login_network_masks").ColumnType(schema.ColumnTypeJSON).Description("The subnet mask that indicates the IP addresses from which logon to the Alibaba Cloud Management Console is allowed. This parameter applies to password-based logon and single sign-on (SSO). However, this parameter does not apply to API calls that are authenticated based on AccessKey pairs. May be more than one CIDR range. If empty then login is allowed from any source.").
			Extractor(column_value_extractor.StructSelector("LoginProfilePreference.LoginNetworkMasks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_user_to_manage_access_keys").ColumnType(schema.ColumnTypeBool).Description("Indicates whether RAM users can manage their AccessKey pairs.").
			Extractor(column_value_extractor.StructSelector("AccessKeyPreference.AllowUserToManageAccessKeys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_user_to_manage_mfa_devices").ColumnType(schema.ColumnTypeBool).Description("Indicates whether RAM users can manage their MFA devices.").
			Extractor(column_value_extractor.StructSelector("MFAPreference.AllowUserToManageMFADevices")).Build(),
	}
}

func (x *TableAlicloudRamSecurityPreferenceGenerator) GetSubTables() []*schema.Table {
	return nil
}
