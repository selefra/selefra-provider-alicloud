package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator{}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetTableName() string {
	return "alicloud_ecs_instance_metric_cpu_utilization_hourly"
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			data := task.ParentRawResult.(ecs.Instance)
			_, err := listCMMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "HOURLY", "acs_ecs_dashboard", "CPUUtilization", "instanceId", data.InstanceId)
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance.").
			Extractor(column_value_extractor.StructSelector("DimensionValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp used for the data point.").Build(),
	}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator) GetSubTables() []*schema.Table {
	return nil
}
