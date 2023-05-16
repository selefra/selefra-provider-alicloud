package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudVpcDhcpOptionsSetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcDhcpOptionsSetGenerator{}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetTableName() string {
	return "alicloud_vpc_dhcp_options_set"
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := vpc.CreateListDhcpOptionsSetsRequest()
			request.Scheme = "https"
			request.MaxResults = requests.NewInteger(100)

			pageLeft := true
			for pageLeft {
				response, err := client.ListDhcpOptionsSets(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, dhcpOptionSet := range response.DhcpOptionsSets {
					resultChannel <- dhcpOptionSet

				}

				if response.NextToken != "" {
					request.NextToken = response.NextToken
				} else {
					pageLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getVpcDhcpOptionSetAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	var id, region string
	region = taskClient.(*alicloud_client.AliCloudClient).Region

	switch item := result.(type) {
	case *vpc.GetDhcpOptionsSetResponse:
		id = item.DhcpOptionsSetId
	case vpc.DhcpOptionsSet:
		id = item.DhcpOptionsSetId
	}

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":dhcpoptionset/" + id}

	return akas, nil
}
func getVpcDhcpOptionsSet(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id string
	if result != nil {
		id = result.(vpc.DhcpOptionsSet).DhcpOptionsSetId
	}

	request := vpc.CreateGetDhcpOptionsSetRequest()
	request.Scheme = "https"
	request.DhcpOptionsSetId = id

	response, err := client.GetDhcpOptionsSet(request)
	if err != nil {

		return nil, nil
	}
	if response.DhcpOptionsSetId != "" {
		return response, nil
	}

	return nil, nil
}
func vpcDhcpOptionsetRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region
	return region, nil
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the DHCP option set.").
			Extractor(column_value_extractor.StructSelector("DhcpOptionsSetName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the DHCP option set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).Description("The root domain.").
			Extractor(column_value_extractor.StructSelector("DhcpOptions.DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Description("The ID of the account to which the DHCP options set belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("DhcpOptionsSetName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getVpcDhcpOptionSetAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcDhcpOptionsetRegion(ctx, clientMeta, taskClient, task, row, column, result)
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
		table_schema_generator.NewColumnBuilder().ColumnName("dhcp_options_set_id").ColumnType(schema.ColumnTypeString).Description("The ID of the DHCP option set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_file_name").ColumnType(schema.ColumnTypeString).Description("The boot file name of DHCP option set.").
			Extractor(column_value_extractor.StructSelector("DhcpOptions.BootFileName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name_servers").ColumnType(schema.ColumnTypeString).Description("The IP addresses of your DNS servers.").
			Extractor(column_value_extractor.StructSelector("DhcpOptions.DomainNameServers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associate_vpc_count").ColumnType(schema.ColumnTypeInt).Description("The number of VPCs associated with DHCP option set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description for the DHCP option set.").
			Extractor(column_value_extractor.StructSelector("DhcpOptionsSetDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tftp_server_name").ColumnType(schema.ColumnTypeString).Description("The tftp server name of the DHCP option set.").
			Extractor(column_value_extractor.StructSelector("DhcpOptions.TFTPServerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associate_vpcs").ColumnType(schema.ColumnTypeJSON).Description("The information of the VPC network that is associated with the DHCP options set.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getVpcDhcpOptionsSet(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAlicloudVpcDhcpOptionsSetGenerator) GetSubTables() []*schema.Table {
	return nil
}
