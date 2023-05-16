package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudEcsNetworkInterfaceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsNetworkInterfaceGenerator{}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetTableName() string {
	return "alicloud_ecs_network_interface"
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeNetworkInterfacesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeNetworkInterfaces(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, eni := range response.NetworkInterfaceSets.NetworkInterfaceSet {
					resultChannel <- eni
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

func ecsEniAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	eni := result.(ecs.NetworkInterfaceSet)
	akas := []string{"acs:ecs:" + eni.ZoneId + ":" + eni.OwnerId + ":eni/" + eni.NetworkInterfaceId}

	return akas, nil
}
func ecsEniTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	eni := result.(ecs.NetworkInterfaceSet)

	title := eni.NetworkInterfaceId

	if len(eni.NetworkInterfaceName) > 0 {
		title = eni.NetworkInterfaceName
	}

	return title, nil
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_id").ColumnType(schema.ColumnTypeString).Description("The ID of the distributor to which the ENI belongs.").
			Extractor(column_value_extractor.StructSelector("ServiceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC to which the ENI belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The zone ID of the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_public_ip_allocation_id").ColumnType(schema.ColumnTypeString).Description("The allocation ID of the EIP.").
			Extractor(column_value_extractor.StructSelector("AssociatedPublicIp.AllocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := ecsEniTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_managed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user is an Alibaba Cloud service or a distributor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interface_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeIp).Description("The private IP address of the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachment").ColumnType(schema.ColumnTypeJSON).Description("Attachments of the ENI").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_sets").ColumnType(schema.ColumnTypeJSON).Description("The private IP addresses of the ENI.").
			Extractor(column_value_extractor.StructSelector("PrivateIpSets.PrivateIpSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the ENI.").
			Extractor(column_value_extractor.StructSelector("NetworkInterfaceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the ENI. Valid values: 'Primary' and 'Secondary'").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the ENI was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Description("The ID of the account that owns the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance to which the ENI is bound.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queue_number").ColumnType(schema.ColumnTypeInt).Description("The number of queues supported by the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_public_ip_address").ColumnType(schema.ColumnTypeIp).Description("The public IP address of the instance.").
			Extractor(column_value_extractor.StructSelector("AssociatedPublicIp.PublicIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_sets").ColumnType(schema.ColumnTypeJSON).Description("The IPv6 addresses assigned to the ENI.").
			Extractor(column_value_extractor.StructSelector("Ipv6Sets.Ipv6Set")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mac_address").ColumnType(schema.ColumnTypeString).Description("The MAC address of the ENI.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the ENI belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_group_ids").ColumnType(schema.ColumnTypeJSON).Description("The IDs of the security groups to which the ENI belongs.").
			Extractor(column_value_extractor.StructSelector("SecurityGroupIds.SecurityGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VSwitch to which the ENI is connected.").
			Extractor(column_value_extractor.StructSelector("VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := ecsEniAka(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAlicloudEcsNetworkInterfaceGenerator) GetSubTables() []*schema.Table {
	return nil
}
