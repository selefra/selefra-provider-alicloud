package tables

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/sethvargo/go-retry"
)

type TableAlicloudRdsInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudRdsInstanceGenerator{}

func (x *TableAlicloudRdsInstanceGenerator) GetTableName() string {
	return "alicloud_rds_instance"
}

func (x *TableAlicloudRdsInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudRdsInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudRdsInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudRdsInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := rds.CreateDescribeDBInstancesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeDBInstances(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range response.Items.DBInstance {
					resultChannel <- i
					count++
				}
				if count >= response.TotalRecordCount {
					break
				}
				request.PageNumber = requests.NewInteger(response.PageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func databaseID(item interface{}) string {
	switch item := item.(type) {
	case rds.DBInstance:
		return item.DBInstanceId
	case rds.DBInstanceAttribute:
		return item.DBInstanceId
	}
	return ""
}
func getRdsInstance(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id string
	if result != nil {
		id = databaseID(result)
	}

	request := rds.CreateDescribeDBInstanceAttributeRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	var response *rds.DescribeDBInstanceAttributeResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeDBInstanceAttribute(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if response.Items.DBInstanceAttribute != nil && len(response.Items.DBInstanceAttribute) > 0 {
		if response.Items.DBInstanceAttribute[0].RegionId == region {
			return response.Items.DBInstanceAttribute[0], nil
		}
	}

	return nil, nil
}
func getRdsInstanceARN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	var region, instanceID string
	switch item := result.(type) {
	case rds.DBInstance:
		region = item.RegionId
		instanceID = item.DBInstanceId
	case rds.DBInstanceAttribute:
		region = item.RegionId
		instanceID = item.DBInstanceId
	}

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return "arn:acs:rds:" + region + ":" + accountID + ":instance/" + instanceID, nil
}
func getRdsInstanceIPArrayList(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id = databaseID(result)

	request := rds.CreateDescribeDBInstanceIPArrayListRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceIPArrayList(request)
	if err != nil {

		return nil, err
	}

	if response.Items.DBInstanceIPArray != nil && len(response.Items.DBInstanceIPArray) > 0 {
		return response, nil
	}

	return nil, nil
}
func getRdsInstanceParameters(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id = databaseID(result)

	request := rds.CreateDescribeParametersRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeParameters(request)
	if err != nil {

		return nil, err
	}
	return response, nil
}
func getRdsTags(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := rds.CreateDescribeTagsRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(result)
	response, err := client.DescribeTags(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {

			return nil, nil
		}

		return nil, err
	}

	if response != nil {
		return response, nil
	}

	return nil, nil
}
func getSSLDetails(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id string
	if result != nil {
		id = databaseID(result)
	}

	request := rds.CreateDescribeDBInstanceSSLRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceSSL(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {

			return nil, nil
		}

		return nil, err
	}
	if len(response.SSLExpireTime) > 0 {
		return "Enabled", nil
	}
	return "Disabled", nil
}
func getSqlCollectorPolicy(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := rds.CreateDescribeSQLCollectorPolicyRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(result)
	response, err := client.DescribeSQLCollectorPolicy(request)
	if err != nil {

		return nil, err
	}
	return response, nil
}
func getSqlCollectorRetention(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	region := taskClient.(*alicloud_client.AliCloudClient).Region

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := rds.CreateDescribeSQLCollectorRetentionRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(result)
	response, err := client.DescribeSQLCollectorRetention(request)
	if err != nil {

		return nil, err
	}
	return response, nil
}
func getTDEDetails(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := alicloud_client.RDSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	var id string
	if result != nil {
		id = databaseID(result)
	}

	request := rds.CreateDescribeDBInstanceTDERequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceTDE(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" || serverErr.ErrorCode() == "InstanceEngineType.NotSupport" || serverErr.ErrorCode() == "InvaildEngineInRegion.ValueNotSupported" {

			return nil, nil
		}

		return nil, err
	}
	return response, nil
}

func (x *TableAlicloudRdsInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudRdsInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("max_connections").ColumnType(schema.ColumnTypeInt).Description("The maximum number of concurrent connections that are allowed by the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MaxConnections")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tde_status").ColumnType(schema.ColumnTypeString).Description("The TDE status at the instance level. Valid values: Enable | Disable.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getTDEDetails(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TDEStatus")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_ips").ColumnType(schema.ColumnTypeJSON).Description("An array that consists of IP addresses in the IP address whitelist.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstanceIPArrayList(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Items.DBInstanceIPArray")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_net_type").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the VPC belongs.").
			Extractor(column_value_extractor.StructSelector("DBInstanceNetType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_max_iops").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryMaxIOPS")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lock_reason").ColumnType(schema.ColumnTypeString).Description("The reason why the instance is locked.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the primary instance to which the instance is attached. If this parameter is not returned, the instance is a primary instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryClass")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collation").ColumnType(schema.ColumnTypeString).Description("The character set collation of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Collation")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_max_quantity").ColumnType(schema.ColumnTypeInt).Description("The maximum number of databases that can be created on the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBMaxQuantity")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multiple_temp_upgrade").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MultipleTempUpgrade")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ins_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("InsId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRdsTags(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the RDS instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRdsInstanceARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The creation time of the Instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("increment_source_db_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance from which incremental data comes. The incremental data of a disaster recovery or read-only instance comes from its primary instance. If this parameter is not returned, the instance is a primary instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("IncrementSourceDBInstanceId")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("readonly_db_instance_ids").ColumnType(schema.ColumnTypeJSON).Description("An array that consists of the IDs of the read-only instances attached to the primary instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ReadOnlyDBInstanceIds")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_network_type").ColumnType(schema.ColumnTypeString).Description("The network type of the instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_disk_used").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBInstanceDiskUsed")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_cpu").ColumnType(schema.ColumnTypeString).Description("The number of CPUs that are configured for the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBInstanceCPU")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_memory").ColumnType(schema.ColumnTypeFloat).Description("The memory capacity of the instance. Unit: MB.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBInstanceMemory")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("IPType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_memory").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryMemory")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_string").ColumnType(schema.ColumnTypeString).Description("The internal endpoint of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ConnectionString")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_upgrade_minor_version").ColumnType(schema.ColumnTypeString).Description("The method that is used to update the minor engine version of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AutoUpgradeMinorVersion")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_zone").ColumnType(schema.ColumnTypeString).Description("The time zone of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TimeZone")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the instances belong.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_status").ColumnType(schema.ColumnTypeString).Description("The status of the instances").
			Extractor(column_value_extractor.StructSelector("DBInstanceStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The ID of the zone to which the instances belong.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_kernel_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("LatestKernelVersion")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_value").ColumnType(schema.ColumnTypeString).Description("The availability status of the instance. Unit: %.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AvailabilityValue")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_ips_src").ColumnType(schema.ColumnTypeJSON).Description("An array that consists of IP details.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstanceIPArrayList(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Items.DBInstanceIPArray")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_collector_retention").ColumnType(schema.ColumnTypeInt).Description("The log backup retention duration that is allowed by the SQL explorer feature on the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getSqlCollectorRetention(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ConfigValue")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_id").ColumnType(schema.ColumnTypeString).Description("The ID of the region to which the instances belong.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRdsInstanceARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintain_time").ColumnType(schema.ColumnTypeString).Description("The maintenance window of the instance. The maintenance window is displayed in UTC+8 in the ApsaraDB RDS console.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MaintainTime")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_time_end").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeTimeEnd")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The ID of the vSwitch associated with the specified VPC.").
			Extractor(column_value_extractor.StructSelector("VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_features").ColumnType(schema.ColumnTypeString).Description("An array that consists of advanced features. The advanced features are separated by commas (,). This parameter is supported only for instances that run SQL Server.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AdvancedFeatures")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).Description("The list of running parameters for the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRdsInstanceParameters(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category").ColumnType(schema.ColumnTypeString).Description("The RDS edition of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeTimestamp).Description("Instance expire time").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_mode").ColumnType(schema.ColumnTypeString).Description("The connection mode of the instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_time_start").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeTimeStart")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_max_connections").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryMaxConnections")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_cpu").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryCpu")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
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
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_storage_type").ColumnType(schema.ColumnTypeString).Description("The type of storage media that is used by the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBInstanceStorageType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_db_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the temporary instance that is attached to the instance if a temporary instance is deployed.").
			Extractor(column_value_extractor.StructSelector("TempDBInstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_type").ColumnType(schema.ColumnTypeString).Description("The role of the instances.").
			Extractor(column_value_extractor.StructSelector("DBInstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_upgrade_account_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("SupportUpgradeAccountType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_status").ColumnType(schema.ColumnTypeString).Description("The SSL encryption status of the Instance").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getSSLDetails(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_description").ColumnType(schema.ColumnTypeString).Description("The description of the DB Instance.").
			Extractor(column_value_extractor.StructSelector("DBInstanceDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).Description("The database engine that the instances run.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lock_mode").ColumnType(schema.ColumnTypeString).Description("The lock mode of the instance.").
			Extractor(column_value_extractor.StructSelector("LockMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_iops").ColumnType(schema.ColumnTypeInt).Description("The maximum number of I/O requests that the instance can process per second.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MaxIOPS")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("super_permission_mode").ColumnType(schema.ColumnTypeString).Description("Indicates whether the instance supports superuser accounts, such as the system administrator (SA) account, Active Directory (AD) account, and host account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("SuperPermissionMode")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_cloud_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the cloud instance on which the specified VPC is deployed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pay_type").ColumnType(schema.ColumnTypeString).Description("The billing method of the instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_class").ColumnType(schema.ColumnTypeString).Description("The instance type of the instances.").
			Extractor(column_value_extractor.StructSelector("DBInstanceClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).Description("The version of the database engine that the instances run.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeString).Description("The internal port of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Port")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_ip_mode").ColumnType(schema.ColumnTypeString).Description("The network isolation mode of the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("SecurityIPMode")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC to which the instances belong.").
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guard_db_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the disaster recovery instance that is attached to the instance if a disaster recovery instance is deployed.").
			Extractor(column_value_extractor.StructSelector("GuardDBInstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the single instance to query.").
			Extractor(column_value_extractor.StructSelector("DBInstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin_configuration").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("OriginConfiguration")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_type").ColumnType(schema.ColumnTypeInt).Description("The type of proxy that is enabled on the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ProxyType")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_collector_policy").ColumnType(schema.ColumnTypeJSON).Description("The status of the SQL Explorer (SQL Audit) feature.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getSqlCollectorPolicy(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRdsTags(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("temp_upgrade_recovery_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TempUpgradeRecoveryTime")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("console_version").ColumnType(schema.ColumnTypeString).Description("The type of proxy that is enabled on the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ConsoleVersion")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_create_super_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("SupportCreateSuperAccount")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dispense_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DispenseMode")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_max_quantity").ColumnType(schema.ColumnTypeInt).Description("The maximum number of accounts that can be created on the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountMaxQuantity")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("DBInstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_host_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the dedicated cluster to which the instances belong if the instances are created in a dedicated cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_storage").ColumnType(schema.ColumnTypeInt).Description("The type of storage media that is used by the instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getRdsInstance(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DBInstanceStorage")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableAlicloudRdsInstanceGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlicloudRdsInstanceMetricConnectionsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudRdsInstanceMetricCpuUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudRdsInstanceMetricCpuUtilizationGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudRdsInstanceMetricConnectionsDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudRdsInstanceMetricCpuUtilizationHourlyGenerator{}),
	}
}
