package tables

import (
	"context"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudVpcRouteEntryGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcRouteEntryGenerator{}

func (x *TableAlicloudVpcRouteEntryGenerator) GetTableName() string {
	return "alicloud_vpc_route_entry"
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			routeTable := task.ParentRawResult.(vpc.RouterTableListType)

			request := vpc.CreateDescribeRouteEntryListRequest()
			request.Scheme = "https"
			request.RouteTableId = routeTable.RouteTableId

			response, err := client.DescribeRouteEntryList(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range response.RouteEntrys.RouteEntry {
				resultChannel <- i
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getVpcRouteEntryTurbotData(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	data := result.(vpc.RouteEntry)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	var title string
	var akas []string
	if len(data.RouteEntryId) > 0 {
		akas = []string{"acs:vpc:" + region + ":" + accountID + ":route-entry/" + data.RouteEntryId}
		title = data.RouteEntryName
	} else {
		akas = []string{"acs:vpc:" + region + ":" + accountID + ":route-entry/" + data.RouteTableId}
		if len(data.NextHops.NextHop[0].NextHopId) > 0 {
			title = data.RouteTableId + ":" + data.DestinationCidrBlock + ":" + data.NextHops.NextHop[0].NextHopType + ":" + data.NextHops.NextHop[0].NextHopId
		} else {
			title = data.RouteTableId + ":" + data.DestinationCidrBlock + ":" + data.NextHops.NextHop[0].NextHopType
		}
	}

	turbotData := map[string]interface{}{
		"Akas":   akas,
		"Title":  title,
		"Region": region,
	}

	return turbotData, nil
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the VRouter.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_version").ColumnType(schema.ColumnTypeString).Description("The version of the IP protocol.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hop_oppsite_type").ColumnType(schema.ColumnTypeString).Description("The type of the next hop.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcRouteEntryTurbotData(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_entry_id").ColumnType(schema.ColumnTypeString).Description("The ID of the route entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the route entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeIp).Description("Specifies the private ip address for the route entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hop_type").ColumnType(schema.ColumnTypeString).Description("The type of the next hop.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destination_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The destination Classless Inter-Domain Routing (CIDR) block of the route entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hop_oppsite_region_id").ColumnType(schema.ColumnTypeString).Description("The region where the next hop instance is deployed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_table_id").ColumnType(schema.ColumnTypeString).Description("The ID of the route table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance associated with the next hop.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hop_oppsite_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance associated with the next hop.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hop_region_id").ColumnType(schema.ColumnTypeString).Description("The region where the next hop instance is deployed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcRouteEntryTurbotData(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcRouteEntryTurbotData(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the route entry.").
			Extractor(column_value_extractor.StructSelector("RouteEntryName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the route entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_hops").ColumnType(schema.ColumnTypeJSON).Description("The information about the next hop.").
			Extractor(column_value_extractor.StructSelector("NextHops.NextHop")).Build(),
	}
}

func (x *TableAlicloudVpcRouteEntryGenerator) GetSubTables() []*schema.Table {
	return nil
}
