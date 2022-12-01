package tables

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudCasCertificateGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudCasCertificateGenerator{}

func (x *TableAlicloudCasCertificateGenerator) GetTableName() string {
	return "alicloud_cas_certificate"
}

func (x *TableAlicloudCasCertificateGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudCasCertificateGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudCasCertificateGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudCasCertificateGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*alicloud_client.AliCloudClient).Region

			if _, exists := supportedRegions[region]; !exists {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			client, err := alicloud_client.CasService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := cas.CreateDescribeUserCertificateListRequest()
			request.ShowSize = "50"
			request.CurrentPage = "1"
			request.QueryParams["RegionId"] = region

			count := 0
			for {
				response, err := client.DescribeUserCertificateList(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, i := range response.CertificateList {
					resultChannel <- i
					count++
				}
				if count >= response.TotalCount {
					break
				}
				request.CurrentPage = requests.NewInteger(response.CurrentPage + 1)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func casCertificate(item interface{}) int64 {
	switch item := item.(type) {
	case cas.Certificate:
		return item.Id
	case *cas.DescribeUserCertificateDetailResponse:
		return item.Id
	}
	return 0
}
func getUserCertificate(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	if _, exists := supportedRegions[region]; !exists {
		return nil, nil
	}

	client, err := alicloud_client.CasService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id int64
	if result != nil {
		data := casCertificate(result)
		id = data
	}

	request := cas.CreateDescribeUserCertificateDetailRequest()
	request.CertId = requests.NewInteger(int(id))

	response, err := client.DescribeUserCertificateDetail(request)
	if err != nil {

		return nil, err
	}

	return response, nil
}
func getUserCertificateAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	data := casCertificate(result)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cas:" + region + ":" + accountID + ":certificate/" + strconv.Itoa(int(data))}

	return akas, nil
}
func getUserCertificateRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	return region, nil
}

func (x *TableAlicloudCasCertificateGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudCasCertificateGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("start_date").ColumnType(schema.ColumnTypeTimestamp).Description("The issuance date of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("city").ColumnType(schema.ColumnTypeString).Description("The city where the organization that purchases the certificate is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cert").ColumnType(schema.ColumnTypeString).Description("The certificate content, in PEM format.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getUserCertificate(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getUserCertificateAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeFloat).Description("The ID of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("buy_in_aliyun").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the certificate was purchased from Alibaba Cloud.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).Description("The certificate fingerprint.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org_name").ColumnType(schema.ColumnTypeString).Description("The name of the organization that purchases the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_date").ColumnType(schema.ColumnTypeTimestamp).Description("The expiration date of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer").ColumnType(schema.ColumnTypeString).Description("The certificate authority.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("common").ColumnType(schema.ColumnTypeString).Description("The common name (CN) attribute of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("province").ColumnType(schema.ColumnTypeString).Description("The province where the organization that purchases the certificate is located.").Build(),
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
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getUserCertificateRegion(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expired").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the certificate has expired.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sans").ColumnType(schema.ColumnTypeString).Description("All domain names bound to the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Description("The country where the organization that purchases the certificate is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).Description("The private key of the certificate, in PEM format.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getUserCertificate(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableAlicloudCasCertificateGenerator) GetSubTables() []*schema.Table {
	return nil
}
