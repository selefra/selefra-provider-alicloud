package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudCmsMonitorHostGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudCmsMonitorHostGenerator{}

func (x *TableAlicloudCmsMonitorHostGenerator) GetTableName() string {
	return "alicloud_cms_monitor_host"
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.CmsService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := cms.CreateDescribeMonitoringAgentHostsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeMonitoringAgentHosts(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, host := range response.Hosts.Host {
					resultChannel <- host
					count++
				}
				if count >= response.Total {
					break
				}
				request.PageNumber = requests.NewInteger(response.PageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getCmsMonitoringAgentStatus(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.CmsService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	id := result.(cms.Host).InstanceId

	request := cms.CreateDescribeMonitoringAgentStatusesRequest()
	request.Scheme = "https"
	request.InstanceIds = id

	response, err := client.DescribeMonitoringAgentStatuses(request)
	if err != nil {

		return nil, err
	}

	return response.NodeStatusList.NodeStatus, nil
}
func getCmsMonitoringHostAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(cms.Host)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:cms:" + data.Region + ":" + accountID + ":host/" + data.HostName}

	return akas, nil
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return alicloud_client.BuildRegionList()
	return nil
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("nat_ip").ColumnType(schema.ColumnTypeString).Description("The IP address of the Network Address Translation (NAT) gateway.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).Description("The serial number of the host. A host that is not provided by Alibaba Cloud has a serial number instead of an instance ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("HostName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type_family").ColumnType(schema.ColumnTypeString).Description("The type of the ECS instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_aliyun_host").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the host is provided by Alibaba Cloud.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("eip_id").ColumnType(schema.ColumnTypeString).Description("The ID of the EIP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("eip_address").ColumnType(schema.ColumnTypeString).Description("The elastic IP address (EIP) of the host.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ali_uid").ColumnType(schema.ColumnTypeBigInt).Description("The ID of the Alibaba Cloud account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_group").ColumnType(schema.ColumnTypeString).Description("The IP address of the host.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operating_system").ColumnType(schema.ColumnTypeString).Description("The operating system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_name").ColumnType(schema.ColumnTypeString).Description("The name of the host.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_version").ColumnType(schema.ColumnTypeString).Description("The version of the Cloud Monitor agent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring_agent_status").ColumnType(schema.ColumnTypeJSON).Description("The status of the Cloud Monitor agent.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCmsMonitoringAgentStatus(ctx, clientMeta, taskClient, task, row, column, result)

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
		table_schema_generator.NewColumnBuilder().ColumnName("network_type").ColumnType(schema.ColumnTypeString).Description("The type of the network.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getCmsMonitoringHostAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
	}
}

func (x *TableAlicloudCmsMonitorHostGenerator) GetSubTables() []*schema.Table {
	return nil
}
