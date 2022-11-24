package tables

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudRamPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRamPolicyGenerator{}

func (x *TableAlicloudRamPolicyGenerator) GetTableName() string {
	return "alicloud_ram_policy"
}

func (x *TableAlicloudRamPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRamPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRamPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRamPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ram.CreateListPoliciesRequest()
			request.Scheme = "https"
			request.MaxItems = requests.NewInteger(1000)

			for {

				response, err := client.ListPolicies(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, policy := range response.Policies.Policy {
					resultChannel <- policy

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

type CaseSensitiveValue []string
type Policy struct {
	Id         string     `json:"Id,omitempty"`
	Statements Statements `json:"Statement"`
	Version    string     `json:"Version"`
}
type Principal map[string]interface{}
type Statement struct {
	Action       Value                  `json:"Action,omitempty"`
	Condition    map[string]interface{} `json:"Condition,omitempty"`
	Effect       string                 `json:"Effect"`
	NotAction    Value                  `json:"NotAction,omitempty"`
	NotPrincipal Principal              `json:"NotPrincipal,omitempty"`
	NotResource  CaseSensitiveValue     `json:"NotResource,omitempty"`
	Principal    Principal              `json:"Principal,omitempty"`
	Resource     CaseSensitiveValue     `json:"Resource,omitempty"`
	Sid          string                 `json:"Sid,omitempty"`
}
type Statements []Statement
type Value []string

func getPolicyAkas(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := policyName(result)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ram::" + accountID + ":policy/" + data}, nil
}
func getRAMPolicy(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RAMService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var name, policyType string
	if result != nil {
		i := result.(ram.PolicyInListPolicies)
		name = i.PolicyName
		policyType = i.PolicyType
	}

	request := ram.CreateGetPolicyRequest()
	request.Scheme = "https"
	request.PolicyName = name
	request.PolicyType = policyType
	var response *ram.GetPolicyResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(10, b), func(ctx context.Context) error {
		var err error
		response, err = client.GetPolicy(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling.User" {
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

	if response != nil && len(response.Policy.PolicyName) > 0 {
		return response, nil
	}

	return nil, nil
}
func policyName(item interface{}) string {
	switch item := item.(type) {
	case ram.PolicyInListPolicies:
		return item.PolicyName
	case *ram.GetPolicyResponse:
		return item.Policy.PolicyName
	}
	return ""
}

func (x *TableAlicloudRamPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlicloudRamPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Description("Policy creation date").
			Extractor(column_value_extractor.StructSelector("CreateDate", "Policy.CreateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_date").ColumnType(schema.ColumnTypeTimestamp).Description("Last time when policy got updated ").
			Extractor(column_value_extractor.StructSelector("UpdateDate", "Policy.UpdateDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).Description("Contains the details about the policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DefaultPolicyVersion.PolicyDocument")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getPolicyAkas(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document_std").ColumnType(schema.ColumnTypeJSON).Description("Contains the policy document in a canonical form for easier searching.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRAMPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DefaultPolicyVersion.PolicyDocument")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("PolicyName", "Policy.PolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.Constant("global")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Description("The name of the policy.").
			Extractor(column_value_extractor.StructSelector("PolicyName", "Policy.PolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_type").ColumnType(schema.ColumnTypeString).Description("The type of the policy. Valid values: System and Custom.").
			Extractor(column_value_extractor.StructSelector("PolicyType", "Policy.PolicyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachment_count").ColumnType(schema.ColumnTypeInt).Description("The number of references to the policy.").
			Extractor(column_value_extractor.StructSelector("AttachmentCount", "Policy.AttachmentCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_version").ColumnType(schema.ColumnTypeString).Description("Deafult version of the policy").
			Extractor(column_value_extractor.StructSelector("DefaultVersion", "Policy.DefaultVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The policy description").
			Extractor(column_value_extractor.StructSelector("Description", "Policy.Description")).Build(),
	}
}

func (x *TableAlicloudRamPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
