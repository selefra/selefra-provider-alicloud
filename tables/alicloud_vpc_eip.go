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

type TableAlicloudVpcEipGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcEipGenerator{}

func (x *TableAlicloudVpcEipGenerator) GetTableName() string {
	return "alicloud_vpc_eip"
}

func (x *TableAlicloudVpcEipGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcEipGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcEipGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcEipGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := vpc.CreateDescribeEipAddressesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeEipAddresses(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.EipAddresses.EipAddress {
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

func getVpcEipArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(vpc.EipAddress)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:vpc:" + data.RegionId + ":" + accountID + ":eip/" + data.AllocationId

	return arn, nil
}
func getVpcEipTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {

	eip := result.(vpc.EipAddress)

	if eip.Name != "" {
		return eip.Name, nil
	}

	return eip.AllocationId, nil
}

func (x *TableAlicloudVpcEipGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcEipGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("expired_time").ColumnType(schema.ColumnTypeTimestamp).Description("The expiration time of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bandwidth").ColumnType(schema.ColumnTypeString).Description("The peak bandwidth of the EIP. Unit: Mbit/s.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_locks_reason").ColumnType(schema.ColumnTypeJSON).Description("The reason why the EIP is locked. Valid values: financial security.").
			Extractor(column_value_extractor.StructSelector("OperationLocks.LockReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the EIP was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hd_monitor_status").ColumnType(schema.ColumnTypeString).Description("Indicates whether fine-grained monitoring is enabled for the EIP.").
			Extractor(column_value_extractor.StructSelector("HDMonitorStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_reservation_data").ColumnType(schema.ColumnTypeBool).Description("Indicates whether renewal data is included.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).Description("The type of the instance to which the EIP is bound.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcEipTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_regions").ColumnType(schema.ColumnTypeJSON).Description("The ID of the region to which the EIP belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_region_id").ColumnType(schema.ColumnTypeString).Description("The region ID of the bound resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the EIP.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpcEipArn(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("charge_type").ColumnType(schema.ColumnTypeString).Description("The billing method of the EIP").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("segment_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance with which the contiguous EIP is associated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the EIP.").
			Extractor(column_value_extractor.StructSelector("Descritpion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("isp").ColumnType(schema.ColumnTypeString).Description("The Internet service provider.").
			Extractor(column_value_extractor.StructSelector("ISP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bandwidth_package_bandwidth").ColumnType(schema.ColumnTypeString).Description("The maximum bandwidth of the EIP in Mbit/s.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bandwidth_package_type").ColumnType(schema.ColumnTypeString).Description("The bandwidth value of the EIP Bandwidth Plan to which the EIP is added.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mode").ColumnType(schema.ColumnTypeString).Description("The type of the instance to which you want to bind the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("netmode").ColumnType(schema.ColumnTypeString).Description("The network type of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("second_limited").ColumnType(schema.ColumnTypeBool).Description("Indicates whether level-2 traffic throttling is configured.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance to which the EIP is bound.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeIp).Description("The IP address of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_charge_type").ColumnType(schema.ColumnTypeString).Description("The metering method of the EIP can be one of PayByBandwidth or PayByTraffic.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_managed").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpcEipArn(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("allocation_id").ColumnType(schema.ColumnTypeString).Description("The unique ID of the EIP.").Build(),
	}
}

func (x *TableAlicloudVpcEipGenerator) GetSubTables() []*schema.Table {
	return nil
}
