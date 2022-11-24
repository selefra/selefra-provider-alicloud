package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudRamAccessKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamAccessKeyGenerator{}

func (x *TableAlicloudRamAccessKeyGenerator) GetTableName() string {
	return "alicloud_ram_access_key"
}

func (x *TableAlicloudRamAccessKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamAccessKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamAccessKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamAccessKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			user := task.ParentRawResult.(userInfo)

			request := ram.CreateListAccessKeysRequest()
			request.Scheme = "https"
			request.UserName = user.UserName

			response, err := client.ListAccessKeys(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range response.AccessKeys.AccessKey {
				resultChannel <- accessKeyRow{i.AccessKeyId, i.Status, i.CreateDate, user.UserName}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type accessKeyRow = struct {
	AccessKeyId string
	Status      string
	CreateDate  string
	UserName    string
}

func getAccessKeyArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	i := result.(accessKeyRow)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ram::" + accountID + ":user/" + i.UserName + "/accesskey/" + i.AccessKeyId}

	return akas, nil
}

func (x *TableAlicloudRamAccessKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlicloudRamAccessKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_id").ColumnType(schema.ColumnTypeString).Description("The AccessKey ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the AccessKey pair. Valid values: Active and Inactive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the AccessKey pair was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("AccessKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getAccessKeyArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("Name of the User that the access key belongs to.").Build(),
	}
}

func (x *TableAlicloudRamAccessKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
