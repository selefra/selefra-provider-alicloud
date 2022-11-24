package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudVpcVpnGatewayGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcVpnGatewayGenerator{}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetTableName() string {
	return "alicloud_vpc_vpn_gateway"
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := vpc.CreateDescribeVpnGatewaysRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeVpnGateways(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.VpnGateways.VpnGateway {
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

func getVpcVpnGatewayAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(vpc.VpnGateway)
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":vpngateway/" + data.VpnGatewayId}

	return akas, nil
}
func getVpnGatewayRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	return region, nil
}
func vpcVpnGatewayTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(vpc.VpnGateway)

	title := data.VpnGatewayId

	if len(data.Name) > 0 {
		title = data.Name
	}

	return title, nil
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcVpnGatewayTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_status").ColumnType(schema.ColumnTypeString).Description("The business state of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the VPN gateway was created.").
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag").ColumnType(schema.ColumnTypeString).Description("The tag of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipsec_vpn").ColumnType(schema.ColumnTypeString).Description("Indicates whether the IPsec-VPN feature is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_ip").ColumnType(schema.ColumnTypeIp).Description("The public IP address of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec").ColumnType(schema.ColumnTypeString).Description("The maximum bandwidth of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_vpn").ColumnType(schema.ColumnTypeString).Description("Indicates whether the SSL-VPN feature is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VSwitch to which the VPN gateway belongs.").
			Extractor(column_value_extractor.StructSelector("VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC for which the VPN gateway is created.").
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_gateway_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPN gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_propagate").ColumnType(schema.ColumnTypeBool).Description("Indicates whether auto propagate is enabled, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_bgp").ColumnType(schema.ColumnTypeBool).Description("Indicates whether bgp is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reservation_data").ColumnType(schema.ColumnTypeJSON).Description("A set of reservation details.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpcVpnGatewayAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpnGatewayRegion(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_method").ColumnType(schema.ColumnTypeString).Description("The billing method of the VPN gateway.").
			Extractor(column_value_extractor.StructSelector("ChargeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Description("The creation time of the VPC.").
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_max_connections").ColumnType(schema.ColumnTypeInt).Description("The maximum number of concurrent SSL-VPN connections.").Build(),
	}
}

func (x *TableAlicloudVpcVpnGatewayGenerator) GetSubTables() []*schema.Table {
	return nil
}
