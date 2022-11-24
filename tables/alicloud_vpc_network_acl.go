package tables

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudVpcNetworkAclGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudVpcNetworkAclGenerator{}

func (x *TableAlicloudVpcNetworkAclGenerator) GetTableName() string {
	return "alicloud_vpc_network_acl"
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.VpcService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := vpc.CreateDescribeNetworkAclsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeNetworkAcls(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.NetworkAcls.NetworkAcl {
					resultChannel <- i
					count++
				}
				totalcount, err := strconv.Atoi(response.TotalCount)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				pageNumber, err := strconv.Atoi(response.PageNumber)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				if count >= totalcount {
					break
				}
				request.PageNumber = requests.NewInteger(pageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getNetworkACLAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := networkAclData(result)
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":network-acl/" + data["ID"]}

	return akas, nil
}
func networkAclData(item interface{}) map[string]string {
	data := map[string]string{}
	switch item := item.(type) {
	case vpc.NetworkAcl:
		data["ID"] = item.NetworkAclId
		data["Name"] = item.NetworkAclName
	case vpc.NetworkAclAttribute:
		data["ID"] = item.NetworkAclId
		data["Name"] = item.NetworkAclName
	}
	return data
}
func networkAclRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region
	return region, nil
}
func vpcNetworkACLTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	data := networkAclData(result)

	title := data["ID"]

	if len(data["Name"]) > 0 {
		title = data["Name"]
	}

	return title, nil
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeInt).Description("The ID of the owner of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_id").ColumnType(schema.ColumnTypeString).Description("The name of the region where the resource resides.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the network ACL.").
			Extractor(column_value_extractor.StructSelector("NetworkAclName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC associated with the network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the network ACL was created.").
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).Description("A list of associated resources.").
			Extractor(column_value_extractor.StructSelector("Resources.Resource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_acl_id").ColumnType(schema.ColumnTypeString).Description("The ID of the network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := vpcNetworkACLTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getNetworkACLAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := networkAclRegion(ctx, clientMeta, taskClient, task, row, column, result)
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
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingress_acl_entries").ColumnType(schema.ColumnTypeJSON).Description("A list of inbound rules of the network ACL.").
			Extractor(column_value_extractor.StructSelector("IngressAclEntries.IngressAclEntry")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("egress_acl_entries").ColumnType(schema.ColumnTypeJSON).Description("A list of outbound rules of the network ACL.").
			Extractor(column_value_extractor.StructSelector("EgressAclEntries.EgressAclEntry")).Build(),
	}
}

func (x *TableAlicloudVpcNetworkAclGenerator) GetSubTables() []*schema.Table {
	return nil
}
