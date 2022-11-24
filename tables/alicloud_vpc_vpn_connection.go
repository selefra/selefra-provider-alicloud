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

type TableAlicloudVpcVpnConnectionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcVpnConnectionGenerator{}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetTableName() string {
	return "alicloud_vpc_vpn_connection"
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := vpc.CreateDescribeVpnConnectionsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeVpnConnections(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, vpnConnection := range response.VpnConnections.VpnConnection {
					resultChannel <- vpnConnection
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

func getVpnConnectionAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(vpc.VpnConnection)
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:vpc:" + region + ":" + accountID + ":vpnconnection/" + data.VpnConnectionId}

	return akas, nil
}
func getVpnConnectionRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	return region, nil
}
func vpnConnectionTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(vpc.VpnConnection)

	title := data.VpnConnectionId

	if len(data.Name) > 0 {
		title = data.Name
	}

	return title, nil
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("enable_nat_traversal").ColumnType(schema.ColumnTypeBool).Description("Indicates whether to enable the NAT traversal feature.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.EnableNatTraversal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_connection_id").ColumnType(schema.ColumnTypeString).Description("The ID of the IPsec-VPN connection.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.VpnConnectionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_gateway_id").ColumnType(schema.ColumnTypeString).Description("The ID of the customer gateway.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.CustomerGatewayId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("local_subnet").ColumnType(schema.ColumnTypeCIDR).Description("The CIDR block of the virtual private cloud (VPC).").
			Extractor(column_value_extractor.StructSelector("VpnConnection.LocalSubnet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remote_subnet").ColumnType(schema.ColumnTypeCIDR).Description("The CIDR block of the on-premises data center.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.RemoteSubnet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vco_health_check").ColumnType(schema.ColumnTypeJSON).Description("The health check configurations.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.VcoHealthCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_bgp_config").ColumnType(schema.ColumnTypeJSON).Description("BGP configuration information.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.VpnBgpConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpnConnectionAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the IPsec-VPN connection.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("effect_immediately").ColumnType(schema.ColumnTypeBool).Description("Indicates whether IPsec-VPN negotiations are initiated immediately.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.EffectImmediately")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_dpd").ColumnType(schema.ColumnTypeBool).Description("Indicates whether dead peer detection (DPD) is enabled.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.EnableDpd")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_gateway_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPN gateway.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.VpnGatewayId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipsec_config").ColumnType(schema.ColumnTypeJSON).Description("The configurations for Phase 2 negotiations.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.IpsecConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpnConnectionTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
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
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the IPsec-VPN connection.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the IPsec-VPN connection was created.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ike_config").ColumnType(schema.ColumnTypeJSON).Description("The configurations of Phase 1 negotiations.").
			Extractor(column_value_extractor.StructSelector("VpnConnection.IkeConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpnConnectionRegion(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
	}
}

func (x *TableAlicloudVpcVpnConnectionGenerator) GetSubTables() []*schema.Table {
	return nil
}
