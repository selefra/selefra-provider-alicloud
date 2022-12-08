package tables

import (
	"context"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudCsKubernetesClusterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudCsKubernetesClusterGenerator{}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetTableName() string {
	return "alicloud_cs_kubernetes_cluster"
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			// TODO 2022-12-8 17:25:27 may be wrong?
			//region := alicloud_client.GetDefaultRegions(ctx, clientMeta, taskClient, task)

			client, err := alicloud_client.ContainerService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := cs.CreateDescribeClustersV1Request()
			request.Scheme = "https"
			request.QueryParams["RegionId"] = taskClient.(*alicloud_client.AliCloudClient).Region
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeClustersV1(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				var result map[string]interface{}
				err = json.Unmarshal([]byte(response.GetHttpContentString()), &result)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				clusters := result["clusters"].([]interface{})
				pageInfo := result["page_info"].(map[string]interface{})
				TotalCount := pageInfo["total_count"].(float64)
				PageNumber := pageInfo["page_number"].(float64)
				for _, cluster := range clusters {
					clusterAsMap := cluster.(map[string]interface{})
					resultChannel <- clusterAsMap
					count++
				}
				if count >= int(TotalCount) {
					break
				}
				request.PageNumber = requests.NewInteger(int(PageNumber) + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getCsKubernetesClusterARN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(map[string]interface{})

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:cs:" + data["region_id"].(string) + ":" + accountID + ":cluster/" + data["cluster_id"].(string)

	return arn, nil
}
func getCsKubernetesClusterLog(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.ContainerService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	id := result.(map[string]interface{})["cluster_id"].(string)

	request := cs.CreateDescribeClusterLogsRequest()
	request.Scheme = "https"
	request.ClusterId = id

	response, err := client.DescribeClusterLogs(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	if len(response.GetHttpContentString()) > 0 {
		return response.GetHttpContentString(), nil
	}

	return nil, nil
}
func getCsKubernetesClusterNamespace(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	id := result.(map[string]interface{})["cluster_id"].(string)

	client, err := alicloud_client.ContainerService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, nil
	}

	request := requests.NewCommonRequest()
	request.Scheme = "https"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/k8s/" + id + "/namespaces"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return nil, nil
	}

	return response.GetHttpContentString(), nil
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("service_discovery_types").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("service_discovery_types")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The IDs of VSwitches.").
			Extractor(column_value_extractor.StructSelector("vswitch_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("worker_ram_role_name").ColumnType(schema.ColumnTypeString).Description("The name of the RAM role for worker nodes in the cluster.").
			Extractor(column_value_extractor.StructSelector("worker_ram_role_name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The ID of the zone where the cluster is deployed.").
			Extractor(column_value_extractor.StructSelector("zone_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_healthy").ColumnType(schema.ColumnTypeString).Description("The health status of the cluster.").
			Extractor(column_value_extractor.StructSelector("cluster_healthy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_version").ColumnType(schema.ColumnTypeString).Description("The version of the cluster.").
			Extractor(column_value_extractor.StructSelector("current_version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_loadbalancer_id").ColumnType(schema.ColumnTypeString).Description("The ID of the Server Load Balancer (SLB) instance deployed in the cluster.").
			Extractor(column_value_extractor.StructSelector("external_loadbalancer_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_status").ColumnType(schema.ColumnTypeString).Description("The status of cluster nodes.").
			Extractor(column_value_extractor.StructSelector("node_status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_mode").ColumnType(schema.ColumnTypeString).Description("The network type of the cluster.").
			Extractor(column_value_extractor.StructSelector("network_mode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("parameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_id").ColumnType(schema.ColumnTypeString).Description("The ID of the cluster.").
			Extractor(column_value_extractor.StructSelector("cluster_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The number of nodes in the cluster.").
			Extractor(column_value_extractor.StructSelector("size")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_disk_size").ColumnType(schema.ColumnTypeInt).Description("The size of a data disk.").
			Extractor(column_value_extractor.StructSelector("data_disk_size")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_info").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("maintenance_info")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the cluster.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsKubernetesClusterARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("docker_version").ColumnType(schema.ColumnTypeString).Description("The version of Docker.").
			Extractor(column_value_extractor.StructSelector("docker_version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("next_version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the cluster was updated.").
			Extractor(column_value_extractor.StructSelector("updated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsKubernetesClusterARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The status of the cluster.").
			Extractor(column_value_extractor.StructSelector("state")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outputs").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("outputs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile").ColumnType(schema.ColumnTypeString).Description("The identifier of the cluster.").
			Extractor(column_value_extractor.StructSelector("profile")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the cluster belongs.").
			Extractor(column_value_extractor.StructSelector("resource_group_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC used by the cluster.").
			Extractor(column_value_extractor.StructSelector("vpc_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled_migration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("enabled_migration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("init_version").ColumnType(schema.ColumnTypeString).Description("The initial version of the cluster.").
			Extractor(column_value_extractor.StructSelector("init_version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("need_update_agent").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("need_update_agent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_zone").ColumnType(schema.ColumnTypeString).Description("Indicates whether PrivateZone is enabled for the cluster.").
			Extractor(column_value_extractor.StructSelector("private_zone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_window").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("maintenance_window")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_url").ColumnType(schema.ColumnTypeJSON).Description("The endpoints that are open for connections to the cluster.").
			Extractor(column_value_extractor.StructSelector("master_url")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_namespace").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsKubernetesClusterNamespace(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("capabilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_type").ColumnType(schema.ColumnTypeString).Description("The type of the cluster.").
			Extractor(column_value_extractor.StructSelector("cluster_type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_disk_category").ColumnType(schema.ColumnTypeString).Description("The type of data disks.").
			Extractor(column_value_extractor.StructSelector("data_disk_category")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upgrade_components").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("upgrade_components")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meta_data").ColumnType(schema.ColumnTypeJSON).Description("The metadata of the cluster.").
			Extractor(column_value_extractor.StructSelector("meta_data")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the cluster.").
			Extractor(column_value_extractor.StructSelector("tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("region_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the cluster was created.").
			Extractor(column_value_extractor.StructSelector("created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).Description("The Elastic Compute Service (ECS) instance type of cluster nodes.").
			Extractor(column_value_extractor.StructSelector("instance_type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_cidr").ColumnType(schema.ColumnTypeCIDR).Description("The CIDR block of pods in the cluster.").
			Extractor(column_value_extractor.StructSelector("subnet_cidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_cidr").ColumnType(schema.ColumnTypeCIDR).Description("The CIDR block of VSwitches.").
			Extractor(column_value_extractor.StructSelector("vswitch_cidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeString).Description("Container port in Kubernetes.").
			Extractor(column_value_extractor.StructSelector("port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("swarm_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("swarm_mode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_log").ColumnType(schema.ColumnTypeJSON).Description("The logs of a cluster.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCsKubernetesClusterLog(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the cluster.").
			Extractor(column_value_extractor.StructSelector("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_spec").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("cluster_spec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).Description("Indicates whether deletion protection is enabled for the cluster.").
			Extractor(column_value_extractor.StructSelector("deletion_protection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gw_bridge").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("gw_bridge")).Build(),
	}
}

func (x *TableAlicloudCsKubernetesClusterGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlicloudCsKubernetesClusterNodeGenerator{}),
	}
}
