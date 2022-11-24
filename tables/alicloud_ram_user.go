package tables

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudRamUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamUserGenerator{}

func (x *TableAlicloudRamUserGenerator) GetTableName() string {
	return "alicloud_ram_user"
}

func (x *TableAlicloudRamUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ram.CreateListUsersRequest()
			request.Scheme = "https"
			for {
				response, err := client.ListUsers(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.Users.User {
					resultChannel <- userInfo{i.UserName, i.UserId, i.DisplayName, i.Email, i.MobilePhone, i.Comments, i.CreateDate, i.UpdateDate, ""}
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

type userInfo = struct {
	UserName      string
	UserId        string
	DisplayName   string
	Email         string
	MobilePhone   string
	Comments      string
	CreateDate    string
	UpdateDate    string
	LastLoginDate string
}

func getCsUserPermissions(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	data := result.(userInfo)

	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Scheme = "https"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/permissions/users/" + data.UserId
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if _, ok := err.(*errors.ServerError); ok {

		return nil, err
	}

	return response.GetHttpContentString(), nil
}
func getRAMUser(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var name string
	if result != nil {
		i := result.(userInfo)
		name = i.UserName
	}

	request := ram.CreateGetUserRequest()
	request.Scheme = "https"
	request.UserName = name
	var response *ram.GetUserResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.GetUser(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}

				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	data := response.User
	return userInfo{data.UserName, data.UserId, data.DisplayName, data.Email, data.MobilePhone, data.Comments, data.CreateDate, data.UpdateDate, data.LastLoginDate}, nil
}
func getRAMUserGroups(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(userInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListGroupsForUserRequest()
	request.Scheme = "https"
	request.UserName = data.UserName

	response, err := client.ListGroupsForUser(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response.Groups.Group, nil
}
func getRAMUserMfaDevices(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {

	data := result.(userInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListVirtualMFADevicesRequest()
	request.Scheme = "https"

	response, err := client.ListVirtualMFADevices(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	var items []ram.VirtualMFADeviceInListVirtualMFADevices
	for _, i := range response.VirtualMFADevices.VirtualMFADevice {
		if i.User.UserName == data.UserName {
			items = append(items, i)
		}
	}

	return items, nil
}
func getRAMUserPolicies(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(userInfo)

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ram.CreateListPoliciesForUserRequest()
	request.Scheme = "https"
	request.UserName = data.UserName
	var response *ram.ListPoliciesForUserResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.ListPoliciesForUser(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}

				return nil
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
func getUserArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(userInfo)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return "acs:ram::" + accountID + ":user/" + data.UserName, nil
}
func userMfaStatus(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.([]ram.VirtualMFADeviceInListVirtualMFADevices)

	if len(data) > 0 {
		return true, nil
	}

	return false, nil
}

func (x *TableAlicloudRamUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlicloudRamUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_device_serial_number").ColumnType(schema.ColumnTypeString).Description("The serial number of the MFA device.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRAMUserMfaDevices(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getUserArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The username of the RAM user.").
			Extractor(column_value_extractor.StructSelector("UserName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the RAM user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getUserArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_enabled").ColumnType(schema.ColumnTypeBool).Description("The MFA status of the user").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMUserMfaDevices(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if reflect_util.IsNil(r) {
					return nil, nil
				}

				r, err = userMfaStatus(ctx, clientMeta, taskClient, task, row, column, r)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policy").ColumnType(schema.ColumnTypeJSON).Description("A list of policies attached to a RAM user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMUserPolicies(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Policies.Policy")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mobile_phone").ColumnType(schema.ColumnTypeString).Description("The mobile phone number of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM user was modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_mfa_devices").ColumnType(schema.ColumnTypeJSON).Description("The list of MFA devices.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRAMUserMfaDevices(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The unique ID of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The email address of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_login_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM user last logged on to the console by using the password.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cs_user_permissions").ColumnType(schema.ColumnTypeJSON).Description("User permissions for Container Service Kubernetes clusters.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsUserPermissions(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeJSON).Description("A list of groups attached to the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRAMUserGroups(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("UserName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments").ColumnType(schema.ColumnTypeString).Description("The description of the RAM user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the RAM user was created.").Build(),
	}
}

func (x *TableAlicloudRamUserGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlicloudRamAccessKeyGenerator{}),
	}
}
