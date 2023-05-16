package tables

import (
	"context"
	"encoding/base64"

	ims "github.com/alibabacloud-go/ims-20190815/client"

	"github.com/gocarina/gocsv"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudRamCredentialReportGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamCredentialReportGenerator{}

func (x *TableAlicloudRamCredentialReportGenerator) GetTableName() string {
	return "alicloud_ram_credential_report"
}

func (x *TableAlicloudRamCredentialReportGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamCredentialReportGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamCredentialReportGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamCredentialReportGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.IMSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := &ims.GetCredentialReportRequest{}

			response, err := client.GetCredentialReport(request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			data, err := base64.StdEncoding.DecodeString(*response.Content)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			content := string(data[:])

			rows := []*alicloudRamCredentialReportResult{}
			if err := gocsv.UnmarshalString(content, &rows); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, row := range rows {
				row.GeneratedTime = response.GeneratedTime
				resultChannel <- row
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type alicloudRamCredentialReportResult struct {
	GeneratedTime                   *string `csv:"-"`
	UserName                        *string `csv:"user"`
	UserCreationTime                *string `csv:"user_creation_time"`
	UserLastLogon                   *string `csv:"user_last_logon"`
	PasswordExist                   *string `csv:"password_exist"`
	PasswordActive                  *string `csv:"password_active"`
	PasswordLastChanged             *string `csv:"password_last_changed"`
	PasswordNextRotation            *string `csv:"password_next_rotation"`
	MfaActive                       *string `csv:"mfa_active"`
	AccessKey1Exist                 *string `csv:"access_key_1_exist"`
	AccessKey1Active                *string `csv:"access_key_1_active"`
	AccessKey1LastRotated           *string `csv:"access_key_1_last_rotated"`
	AccessKey1LastUsed              *string `csv:"access_key_1_last_used"`
	AccessKey2Exist                 *string `csv:"access_key_2_exist"`
	AccessKey2Active                *string `csv:"access_key_2_active"`
	AccessKey2LastRotated           *string `csv:"access_key_2_last_rotated"`
	AccessKey2LastUsed              *string `csv:"access_key_2_last_used"`
	AdditionalAccessKey1Exist       *string `csv:"additional_access_key_1_exist"`
	AdditionalAccessKey1Active      *string `csv:"additional_access_key_1_active"`
	AdditionalAccessKey1LastRotated *string `csv:"additional_access_key_1_last_rotated"`
	AdditionalAccessKey1LastUsed    *string `csv:"additional_access_key_1_last_used"`
	AdditionalAccessKey2Exist       *string `csv:"additional_access_key_2_exist"`
	AdditionalAccessKey2Active      *string `csv:"additional_access_key_2_active"`
	AdditionalAccessKey2LastRotated *string `csv:"additional_access_key_2_last_rotated"`
	AdditionalAccessKey2LastUsed    *string `csv:"additional_access_key_2_last_used"`
	AdditionalAccessKey3Exist       *string `csv:"additional_access_key_3_exist"`
	AdditionalAccessKey3Active      *string `csv:"additional_access_key_3_active"`
	AdditionalAccessKey3LastRotated *string `csv:"additional_access_key_3_last_rotated"`
	AdditionalAccessKey3LastUsed    *string `csv:"additional_access_key_3_last_used"`
}

func (x *TableAlicloudRamCredentialReportGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudRamCredentialReportGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user access key is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_3_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have access key, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_3_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user access key is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_last_logon").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the user last logged in to the console.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have access key, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_last_used").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key was most recently used to sign an Alicloud API request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_1_last_used").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key was most recently used to sign an Alicloud API request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the password is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user access key is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_2_last_used").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key was most recently used to sign an Alicloud API request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generated_time").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the credential report has been generated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether multi-factor authentication (MFA) device has been enabled for the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the user is created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_3_last_used").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key was most recently used to sign an Alicloud API request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_last_rotated").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key has been rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_last_used").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key was most recently used to sign an Alicloud API request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_1_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have access key, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_2_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have access key, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_last_rotated").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key has been rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_1_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user access key is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_1_last_rotated").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key has been rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_2_active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user access key is active, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("The email of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have any password for logging in, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_next_rotation").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the password will be rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_exist").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user have access key, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_2_last_rotated").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key has been rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_last_changed").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the password has been updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_access_key_3_last_rotated").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the time when the access key has been rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudRamCredentialReportGenerator) GetSubTables() []*schema.Table {
	return nil
}
