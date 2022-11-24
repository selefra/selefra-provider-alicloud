package tables

import (
	"context"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudKmsSecretGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudKmsSecretGenerator{}

func (x *TableAlicloudKmsSecretGenerator) GetTableName() string {
	return "alicloud_kms_secret"
}

func (x *TableAlicloudKmsSecretGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudKmsSecretGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudKmsSecretGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudKmsSecretGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.KMSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := kms.CreateListSecretsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.ListSecrets(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.SecretList.Secret {
					resultChannel <- &kms.DescribeSecretResponse{
						CreateTime:        i.CreateTime,
						PlannedDeleteTime: i.PlannedDeleteTime,
						SecretName:        i.SecretName,
						UpdateTime:        i.UpdateTime,
						SecretType:        i.SecretType,
						Tags: kms.TagsInDescribeSecret{
							Tag: i.Tags.Tag,
						},
					}

					count++
				}
				if count >= response.TotalCount {
					break
				}
				request.PageNumber = requests.NewInteger(response.PageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func fetchRegionFromArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(*kms.DescribeSecretResponse)

	resourceArn := data.Arn
	region := strings.Split(resourceArn, ":")[2]
	return region, nil
}
func getKmsSecret(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.KMSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var name string
	var response *kms.DescribeSecretResponse
	if result != nil {
		data := result.(*kms.DescribeSecretResponse)
		name = data.SecretName
	}

	request := kms.CreateDescribeSecretRequest()
	request.Scheme = "https"
	request.SecretName = name
	request.FetchTags = "true"

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeSecret(request)
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

	return response, nil
}
func listKmsSecretVersionIds(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.KMSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}
	secretData := result.(*kms.DescribeSecretResponse)
	var response *kms.ListSecretVersionIdsResponse

	request := kms.CreateListSecretVersionIdsRequest()
	request.Scheme = "https"
	request.SecretName = secretData.SecretName
	request.IncludeDeprecated = "true"

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.ListSecretVersionIds(request)
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

	if response.VersionIds.VersionId != nil && len(response.VersionIds.VersionId) > 0 {
		return response.VersionIds, nil
	}

	return nil, nil
}

func (x *TableAlicloudKmsSecretGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudKmsSecretGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("secret_type").ColumnType(schema.ColumnTypeString).Description("The type of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_rotation_date").ColumnType(schema.ColumnTypeTimestamp).Description("Date of last rotation of Secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("planned_delete_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the KMS Secret is planned to delete.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_config").ColumnType(schema.ColumnTypeJSON).Description("The extended configuration of Secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Tags.Tag")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
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
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the secret.").
			Extractor(column_value_extractor.StructSelector("SecretName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the KMS Secret was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_rotation_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date of next rotation of Secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the KMS Secret was modifies.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_ids").ColumnType(schema.ColumnTypeJSON).Description("The list of secret versions.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := listKmsSecretVersionIds(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("VersionId")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Tags.Tag")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("SecretName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if reflect_util.IsNil(r) {
					return nil, nil
				}

				r, err = fetchRegionFromArn(ctx, clientMeta, taskClient, task, row, column, r)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automatic_rotation").ColumnType(schema.ColumnTypeString).Description("Specifies whether automatic key rotation is enabled.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_key_id").ColumnType(schema.ColumnTypeString).Description("The ID of the KMS customer master key (CMK) that is used to encrypt the secret value.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_interval").ColumnType(schema.ColumnTypeString).Description("The rotation perion of Secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getKmsSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Arn")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableAlicloudKmsSecretGenerator) GetSubTables() []*schema.Table {
	return nil
}
