package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudVpcSslVpnServerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcSslVpnServerGenerator{}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetTableName() string {
	return "alicloud_vpc_ssl_vpn_server"
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := vpc.CreateDescribeSslVpnServersRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeSslVpnServers(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.SslVpnServers.SslVpnServer {
					resultChannel <- i
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

type alicloudCommonColumnData struct {
	AccountID string
}

func ecsVpnSslServerTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	sslServer := result.(vpc.SslVpnServer)

	title := sslServer.SslVpnServerId

	if len(sslServer.Name) > 0 {
		title = sslServer.Name
	}
	return title, nil
}
func getCallerIdentity(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {

	client, err := alicloud_client.StsService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	request := sts.CreateGetCallerIdentityRequest()
	request.Scheme = "https"

	callerIdentity, err := client.GetCallerIdentity(request)
	if err != nil {

		return nil, err
	}

	return callerIdentity, nil
}

func getCommonColumns(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {

	getCallerIdentityData, err := getCallerIdentity(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}

	callerIdentity := getCallerIdentityData.(*sts.GetCallerIdentityResponse)
	commonColumnData := &alicloudCommonColumnData{
		AccountID: callerIdentity.AccountId,
	}

	return commonColumnData, nil
}
func getVpnSslServerAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	sslServer := result.(vpc.SslVpnServer)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + sslServer.RegionId + ":" + accountID + ":sslVpnServer/" + sslServer.SslVpnServerId}

	return akas, nil
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_compressed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the transmitted data is compressed.").
			Extractor(column_value_extractor.StructSelector("Compress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := ecsVpnSslServerTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_vpn_server_id").ColumnType(schema.ColumnTypeString).Description("The ID of the SSL-VPN server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connections").ColumnType(schema.ColumnTypeInt).Description("The total number of current connections.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_ip_pool").ColumnType(schema.ColumnTypeCIDR).Description("The client IP address pool.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the SSL-VPN server was created.").
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_multi_factor_auth").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the multi factor authenticaton is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_ip").ColumnType(schema.ColumnTypeIp).Description("The public IP address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proto").ColumnType(schema.ColumnTypeString).Description("The protocol used by the SSL-VPN server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_gateway_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cipher").ColumnType(schema.ColumnTypeString).Description("The encryption algorithm.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpnSslServerAka(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("local_subnet").ColumnType(schema.ColumnTypeString).Description("The CIDR block of the client.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeInt).Description("The port used by the SSL-VPN server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the SSL-VPN server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_connections").ColumnType(schema.ColumnTypeInt).Description("The maximum number of connections.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudVpcSslVpnServerGenerator) GetSubTables() []*schema.Table {
	return nil
}
