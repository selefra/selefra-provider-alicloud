package tables

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudVpcVswitchGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcVswitchGenerator{}

func (x *TableAlicloudVpcVswitchGenerator) GetTableName() string {
	return "alicloud_vpc_vswitch"
}

func (x *TableAlicloudVpcVswitchGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcVswitchGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcVswitchGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcVswitchGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := vpc.CreateDescribeVSwitchesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeVSwitches(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.VSwitches.VSwitch {
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

func getVSwitchAttributes(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}
	request := vpc.CreateDescribeVSwitchAttributesRequest()
	request.Scheme = "https"
	i := result.(vpc.VSwitch)
	request.VSwitchId = i.VSwitchId
	response, err := client.DescribeVSwitchAttributes(request)
	if err != nil {

		return nil, err
	}
	return response, nil
}
func vswitchAkas(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	i := result.(vpc.VSwitch)
	return []string{"acs:vswitch:" + i.ZoneId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vswitch/" + i.VSwitchId}, nil
}
func vswitchTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	i := result.(vpc.VSwitch)

	title := i.VSwitchId
	if len(i.VSwitchName) > 0 {
		title = i.VSwitchName
	}

	return title, nil
}

func (x *TableAlicloudVpcVswitchGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcVswitchGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the VPC belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Description("The ID of the owner of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC to which the VSwitch belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The IPv4 CIDR block of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_ip_address_count").ColumnType(schema.ColumnTypeInt).Description("The number of available IP addresses in the VSwitch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("share_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The IPv6 CIDR block of the VPC.").
			Extractor(column_value_extractor.StructSelector("Ipv6CidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The zone to which the VSwitch belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Description("True if the VPC is the default VPC in the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_acl_id").ColumnType(schema.ColumnTypeString).Description("A list of IDs of NAT Gateways.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the VPC.").
			Extractor(column_value_extractor.StructSelector("VSwitchName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The unique ID of the VPC.").
			Extractor(column_value_extractor.StructSelector("VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The creation time of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_table").ColumnType(schema.ColumnTypeJSON).Description("Details of the route table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_resources").ColumnType(schema.ColumnTypeJSON).Description("The list of resources in the VSwitch.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVSwitchAttributes(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("CloudResourceSetType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vswitchTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vswitchAkas(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAlicloudVpcVswitchGenerator) GetSubTables() []*schema.Table {
	return nil
}
