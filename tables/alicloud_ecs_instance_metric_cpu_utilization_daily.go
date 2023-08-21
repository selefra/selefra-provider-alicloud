package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator{}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetTableName() string {
	return "alicloud_ecs_instance_metric_cpu_utilization_daily"
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			data := task.ParentRawResult.(ecs.Instance)
			_, err := listCMMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "DAILY", "acs_ecs_dashboard", "CPUUtilization", "instanceId", data.InstanceId)
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance.").
			Extractor(column_value_extractor.StructSelector("DimensionValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alicloud_ecs_instance_selefra_id").ColumnType(schema.ColumnTypeString).Description("fk to alicloud_ecs_instance.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator) GetSubTables() []*schema.Table {
	return nil
}
