package tables

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudEcsDiskMetricReadIopsDailyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsDiskMetricReadIopsDailyGenerator{}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetTableName() string {
	return "alicloud_ecs_disk_metric_read_iops_daily"
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			data := task.ParentRawResult.(ecs.Instance)
			_, err := listCMMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "DAILY", "acs_ecs_dashboard", "DiskReadIOPS", "instanceId", data.InstanceId)
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

type CMMetricRow struct {
	DimensionName string

	DimensionValue string

	Namespace string

	MetricName string

	Average float64

	Maximum float64

	Minimum float64

	Timestamp string
}

func formatTime(timestamp float64) string {
	timeInSec := math.Floor(timestamp / 1000)
	unixTimestamp := time.Unix(int64(timeInSec), 0)
	timestampRFC3339Format := unixTimestamp.Format(time.RFC3339)
	return timestampRFC3339Format
}
func getCMPeriodForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":

		return "86400"
	case "HOURLY":

		return "3600"
	}

	return "300"
}
func getCMStartDateForGranularity(granularity string) string {
	str := "2006-01-02T15:04:05Z"
	switch strings.ToUpper(granularity) {
	case "DAILY":

		return time.Now().AddDate(0, 0, -30).Format(str)
	case "HOURLY":

		return time.Now().AddDate(0, 0, -30).Format(str)
	}

	return time.Now().AddDate(0, 0, -5).Format(str)
}
func getCustomError(errorMessage string) error {
	return errors.NewServerError(500, errorMessage, "")
}
func listCMMetricStatistics(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any, granularity string, namespace string, metricName string, dimensionName string, dimensionValue string) (*cms.DescribeMetricListResponse, error) {

	client, err := alicloud_client.CmsService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}
	request := cms.CreateDescribeMetricListRequest()
	metricDimension := "[{\"" + dimensionName + "\": \"" + dimensionValue + "\"}]"

	request.MetricName = metricName
	request.StartTime = getCMStartDateForGranularity(granularity)
	request.EndTime = time.Now().Format("2006-01-02T15:04:05Z")
	request.Namespace = namespace
	request.Period = getCMPeriodForGranularity(granularity)
	request.Dimensions = metricDimension
	var stats *cms.DescribeMetricListResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		stats, err = client.DescribeMetricList(request)
		if err != nil || stats.Datapoints == "" {

			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				return err
			}

			if stats.Datapoints == "" && !stats.Success {
				err = getCustomError(fmt.Sprint(stats))
				return retry.RetryableError(err)
			}

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	if stats.Datapoints == "" {
		return nil, nil
	}

	err = json.Unmarshal([]byte(stats.Datapoints), &results)
	if err != nil {
		return nil, err
	}
	for _, pointValue := range results {
		resultChannel <- &CMMetricRow{
			DimensionName:  dimensionName,
			DimensionValue: pointValue[dimensionName].(string),
			Namespace:      namespace,
			MetricName:     metricName,
			Average:        pointValue["Average"].(float64),
			Maximum:        pointValue["Maximum"].(float64),
			Minimum:        pointValue["Minimum"].(float64),
			Timestamp:      formatTime(pointValue["timestamp"].(float64)),
		}
	}

	if stats.NextToken != "" {
		request.NextToken = stats.NextToken
	}

	return nil, nil
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").
			Extractor(column_value_extractor.StructSelector("DimensionValue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alicloud_ecs_instance_selefra_id").ColumnType(schema.ColumnTypeString).Description("fk to alicloud_ecs_instance.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAlicloudEcsDiskMetricReadIopsDailyGenerator) GetSubTables() []*schema.Table {
	return nil
}
