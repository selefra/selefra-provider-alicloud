package tables

import (
	"context"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudCsKubernetesClusterNodeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudCsKubernetesClusterNodeGenerator{}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetTableName() string {
	return "alicloud_cs_kubernetes_cluster_node"
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ContainerService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			clusterId := task.ParentRawResult.(map[string]interface{})["cluster_id"].(string)
			request := cs.CreateDescribeClusterNodesRequest()
			request.Scheme = "https"
			request.ClusterId = clusterId

			response, err := client.DescribeClusterNodes(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, node := range response.Nodes {
				resultChannel <- &NodeInfo{
					ClusterId: clusterId,
					Node:      node,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type NodeInfo struct {
	ClusterId string
	cs.Node
}

func clusterNodeRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	nodeName := result.(*NodeInfo).NodeName

	return strings.Split(nodeName, ".")[0], nil
}
func getCsKubernetesClusterNodeAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	nodeName := result.(*NodeInfo).NodeName

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cs:" + strings.Split(nodeName, ".")[0] + ":" + accountID + ":node/" + nodeName}

	return akas, nil
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).Description("The instance type of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_name").ColumnType(schema.ColumnTypeString).Description("The name of the node in the ACK cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the node was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expired_time").ColumnType(schema.ColumnTypeTimestamp).Description("The expiration time of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_name").ColumnType(schema.ColumnTypeString).Description("The name of the node. This name contains the ID of the cluster to which the node is deployed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_role").ColumnType(schema.ColumnTypeString).Description("The role of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_status").ColumnType(schema.ColumnTypeString).Description("The state of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeString).Description("Indicates how the nodes in the node pool were initialized. The nodes can be manually created or created by using Resource Orchestration Service (ROS).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeString).Description("The IP address of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_name").ColumnType(schema.ColumnTypeString).Description("The name of the host.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := clusterNodeRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nodepool_id").ColumnType(schema.ColumnTypeString).Description("The ID of the node pool.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_message").ColumnType(schema.ColumnTypeString).Description("The error message generated when the node was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_id").ColumnType(schema.ColumnTypeString).Description("The ID of the cluster that the node pool belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_aliyun_node").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the instance is provided by Alibaba Cloud.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the ECS instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_charge_type").ColumnType(schema.ColumnTypeString).Description("The billing method of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type_family").ColumnType(schema.ColumnTypeString).Description("The ECS instance family of the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).Description("The ID of the system image that is used by the node.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The states of the nodes in the node pool.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_status").ColumnType(schema.ColumnTypeString).Description("Indicates whether the node is ready in the ACK cluster. Valid values: true, false.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("NodeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsKubernetesClusterNodeAka(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alicloud_cs_kubernetes_cluster_selefra_id").ColumnType(schema.ColumnTypeString).Description("fk to alicloud_cs_kubernetes_cluster.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAlicloudCsKubernetesClusterNodeGenerator) GetSubTables() []*schema.Table {
	return nil
}
