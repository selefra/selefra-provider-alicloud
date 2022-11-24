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

type TableAlicloudVpcNatGatewayGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcNatGatewayGenerator{}

func (x *TableAlicloudVpcNatGatewayGenerator) GetTableName() string {
	return "alicloud_vpc_nat_gateway"
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := vpc.CreateDescribeNatGatewaysRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeNatGateways(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.NatGateways.NatGateway {
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

func getVpcNatGatewayAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	ngw := result.(vpc.NatGateway)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + ngw.RegionId + ":" + accountID + ":natgateway/" + ngw.NatGatewayId}

	return akas, nil
}
func vpcNatGatewayTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := result.(vpc.NatGateway)

	title := data.NatGatewayId

	if len(data.Name) > 0 {
		title = data.Name
	}

	return title, nil
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("auto_pay").ColumnType(schema.ColumnTypeBool).Description("Indicates whether auto pay is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ecs_metric_enabled").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the traffic monitoring feature is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the virtual private cloud (VPC) to which the NAT gateway belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snat_table_ids").ColumnType(schema.ColumnTypeJSON).Description("The ID of the SNAT table for the NAT gateway.").
			Extractor(column_value_extractor.StructSelector("SnatTableIds.SnatTableId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nat_type").ColumnType(schema.ColumnTypeString).Description("The type of the NAT gateway. Valid values: 'Normal' and 'Enhanced'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).Description("Indicates whether deletion protection is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcNatGatewayTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_status").ColumnType(schema.ColumnTypeString).Description("The status of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forward_table_ids").ColumnType(schema.ColumnTypeJSON).Description("The ID of the Destination Network Address Translation (DNAT) table.").
			Extractor(column_value_extractor.StructSelector("ForwardTableIds.ForwardTableId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpcNatGatewayAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_method").ColumnType(schema.ColumnTypeString).Description("The billing method of the NAT gateway.").
			Extractor(column_value_extractor.StructSelector("InstanceChargeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the NAT gateway was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The state of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expired_ime").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the NAT gateway expires.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_charge_type").ColumnType(schema.ColumnTypeString).Description("The billing method of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec").ColumnType(schema.ColumnTypeString).Description("The size of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_lists").ColumnType(schema.ColumnTypeJSON).Description("The elastic IP address (EIP) that is associated with the NAT gateway.").
			Extractor(column_value_extractor.StructSelector("IpLists.IpList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nat_gateway_private_info").ColumnType(schema.ColumnTypeJSON).Description("The information of the virtual private cloud (VPC) to which the enhanced NAT gateway belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nat_gateway_id").ColumnType(schema.ColumnTypeString).Description("The ID of the NAT gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableAlicloudVpcNatGatewayGenerator) GetSubTables() []*schema.Table {
	return nil
}
