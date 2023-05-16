package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudRamGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamGroupGenerator{}

func (x *TableAlicloudRamGroupGenerator) GetTableName() string {
	return "alicloud_ram_group"
}

func (x *TableAlicloudRamGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ram.CreateListGroupsRequest()
			request.Scheme = "https"

			for {
				response, err := client.ListGroups(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.Groups.Group {
					resultChannel <- groupInfo{i.GroupName, i.Comments, i.CreateDate, i.UpdateDate}
				}
				if !response.IsTruncated {
					break
				}
				request.Marker = response.Marker
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type groupInfo = struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}

func getGroupArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(groupInfo)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return "acs:ram::" + accountID + ":group/" + data.GroupName, nil
}
func getRAMGroupPolicies(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(groupInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListPoliciesForGroupRequest()
	request.Scheme = "https"
	request.GroupName = data.GroupName

	response, err := client.ListPoliciesForGroup(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response, nil
}
func getRAMGroupUsers(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(groupInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListUsersForGroupRequest()
	request.Scheme = "https"
	request.GroupName = data.GroupName

	response, err := client.ListUsersForGroup(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response.Users.User, nil
}

func (x *TableAlicloudRamGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudRamGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM user group was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM user group was modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).Description("A list of users in the group.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRAMGroupUsers(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the RAM user group.").
			Extractor(column_value_extractor.StructSelector("GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the RAM user group.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGroupArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments").ColumnType(schema.ColumnTypeString).Description("The description of the RAM user group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policy").ColumnType(schema.ColumnTypeJSON).Description("A list of policies attached to a RAM user group.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMGroupPolicies(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Policies.Policy")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getGroupArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudRamGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
