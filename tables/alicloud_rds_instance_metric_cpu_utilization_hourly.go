package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator{}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetTableName() string {
	return "alicloud_rds_instance_metric_cpu_utilization_hourly"
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			data := task.ParentRawResult.(rds.DBInstance)
			_, err := listCMMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "HOURLY", "acs_rds_dashboard", "CpuUsage", "instanceId", data.DBInstanceId)
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the single instance to query.").
			Extractor(column_value_extractor.StructSelector("DimensionValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alicloud_rds_instance_selefra_id").ColumnType(schema.ColumnTypeString).Description("fk to alicloud_rds_instance.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator) GetSubTables() []*schema.Table {
	return nil
}
