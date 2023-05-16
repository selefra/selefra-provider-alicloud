package tables

import (
	"context"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudVpcGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcGenerator{}

func (x *TableAlicloudVpcGenerator) GetTableName() string {
	return "alicloud_vpc"
}

func (x *TableAlicloudVpcGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := vpc.CreateDescribeVpcsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeVpcs(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.Vpcs.Vpc {
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

func getVpcAttributes(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}
	request := vpc.CreateDescribeVpcAttributeRequest()
	request.Scheme = "https"
	i := result.(vpc.Vpc)
	request.VpcId = i.VpcId

	var response *vpc.DescribeVpcAttributeResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeVpcAttribute(request)
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
func vpcArn(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	i := result.(vpc.Vpc)
	return "acs:vpc:" + i.RegionId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vpc/" + i.VpcId, nil
}
func vpcTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	i := result.(vpc.Vpc)

	title := i.VpcId
	if len(i.VpcName) > 0 {
		title = i.VpcName
	}

	return title, nil
}

func (x *TableAlicloudVpcGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the VPC.").
			Extractor(column_value_extractor.StructSelector("VpcName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vrouter_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VRouter.").
			Extractor(column_value_extractor.StructSelector("VRouterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_cens").ColumnType(schema.ColumnTypeJSON).Description("The list of Cloud Enterprise Network (CEN) instances to which the VPC is attached. No value is returned if the VPC is not attached to any CEN instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcAttributes(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AsssociatedCens")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_blocks").ColumnType(schema.ColumnTypeJSON).Description("The IPv6 CIDR blocks of the VPC.").
			Extractor(column_value_extractor.StructSelector("Ipv6CidrBlocks.Ipv6CidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_acl_num").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classic_link_enabled").ColumnType(schema.ColumnTypeBool).Description("True if the ClassicLink function is enabled.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcAttributes(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of VSwitches in the VPC.").
			Extractor(column_value_extractor.StructSelector("VSwitchIds.VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cen_status").ColumnType(schema.ColumnTypeString).Description("Indicates whether the VPC is attached to any Cloud Enterprise Network (CEN) instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_advanced_feature").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The unique ID of the VPC.").
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The creation time of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The IPv6 CIDR block of the VPC.").
			Extractor(column_value_extractor.StructSelector("Ipv6CidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_cidrs").ColumnType(schema.ColumnTypeJSON).Description("A list of user CIDRs.").
			Extractor(column_value_extractor.StructSelector("UserCidrs.UserCidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dhcp_options_set_id").ColumnType(schema.ColumnTypeString).Description("The ID of the DHCP options set associated to vpc.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dhcp_options_set_status").ColumnType(schema.ColumnTypeString).Description("The status of the VPC network that is associated with the DHCP options set. Valid values: InUse and Pending").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_resources").ColumnType(schema.ColumnTypeJSON).Description("The list of resources in the VPC.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcAttributes(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("CloudResourceSetType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_cidr_blocks").ColumnType(schema.ColumnTypeJSON).Description("A list of secondary IPv4 CIDR blocks of the VPC.").
			Extractor(column_value_extractor.StructSelector("SecondaryCidrBlocks.SecondaryCidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_resource").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nat_gateway_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of IDs of NAT Gateways.").
			Extractor(column_value_extractor.StructSelector("NatGatewayIds.NatGatewayIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_table_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of IDs of route tables.").
			Extractor(column_value_extractor.StructSelector("RouterTableIds.RouterTableIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the VPC.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcArn(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The IPv4 CIDR block of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Description("True if the VPC is the default VPC in the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the VPC belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Description("The ID of the owner of the VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudVpcGenerator) GetSubTables() []*schema.Table {
	return nil
}
