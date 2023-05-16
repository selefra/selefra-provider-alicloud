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

type TableAlicloudRamRoleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamRoleGenerator{}

func (x *TableAlicloudRamRoleGenerator) GetTableName() string {
	return "alicloud_ram_role"
}

func (x *TableAlicloudRamRoleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamRoleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamRoleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamRoleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := ram.CreateListRolesRequest()
			request.Scheme = "https"

			for {
				response, err := client.ListRoles(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.Roles.Role {
					resultChannel <- roleInfo{i.RoleId, i.RoleName, i.Arn, i.Description, "", i.CreateDate, i.UpdateDate, i.MaxSessionDuration}
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

type roleInfo = struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
	MaxSessionDuration       int64
}

func getRAMRole(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var name string
	if result != nil {
		i := result.(roleInfo)
		name = i.RoleName
	}

	request := ram.CreateGetRoleRequest()
	request.Scheme = "https"
	request.RoleName = name

	response, err := client.GetRole(request)
	if err != nil {

		return nil, err
	}

	data := response.Role
	return roleInfo{data.RoleId, data.RoleName, data.Arn, data.Description, data.AssumeRolePolicyDocument, data.CreateDate, data.UpdateDate, data.MaxSessionDuration}, nil
}
func getRAMRolePolicies(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(roleInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListPoliciesForRoleRequest()
	request.Scheme = "https"
	request.RoleName = data.RoleName

	response, err := client.ListPoliciesForRole(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response, nil
}

func (x *TableAlicloudRamRoleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudRamRoleGenerator) GetColumns() []*schema.Column {
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
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the RAM role.").
			Extractor(column_value_extractor.StructSelector("RoleName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_session_duration").ColumnType(schema.ColumnTypeInt).Description("The maximum session duration of the RAM role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM role was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM role was modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assume_role_policy_document").ColumnType(schema.ColumnTypeJSON).Description("The content of the policy that specifies one or more entities entrusted to assume the RAM role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMRole(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AssumeRolePolicyDocument")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the RAM role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assume_role_policy_document_std").ColumnType(schema.ColumnTypeJSON).Description("The standard content of the policy that specifies one or more entities entrusted to assume the RAM role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMRole(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AssumeRolePolicyDocument")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("RoleName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the RAM role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_id").ColumnType(schema.ColumnTypeString).Description("The ID of the RAM role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policy").ColumnType(schema.ColumnTypeJSON).Description("A list of policies attached to a RAM role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMRolePolicies(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Policies.Policy")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudRamRoleGenerator) GetSubTables() []*schema.Table {
	return nil
}
