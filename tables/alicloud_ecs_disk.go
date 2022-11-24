package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudEcsDiskGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsDiskGenerator{}

func (x *TableAlicloudEcsDiskGenerator) GetTableName() string {
	return "alicloud_ecs_disk"
}

func (x *TableAlicloudEcsDiskGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsDiskGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsDiskGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsDiskGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeDisksRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeDisks(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, disk := range response.Disks.Disk {
					resultChannel <- disk
					count++
				}
				if count >= response.TotalCount {
					break
				}
				request.PageNumber = requests.NewInteger(response.PageNumber + 1)
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func ecsDiskTitle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	disk := result.(ecs.Disk)

	title := disk.DiskId

	if len(disk.DiskName) > 0 {
		title = disk.DiskName
	}

	return title, nil
}
func getEcsDiskARN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	disk := result.(ecs.Disk)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + disk.RegionId + ":" + accountID + ":disk/" + disk.DiskId

	return arn, nil
}
func getEcsDiskAutoSnapshotPolicy(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	disk := result.(ecs.Disk)

	client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ecs.CreateDescribeAutoSnapshotPolicyExRequest()
	request.Scheme = "https"
	request.AutoSnapshotPolicyId = disk.AutoSnapshotPolicyId

	response, err := client.DescribeAutoSnapshotPolicyEx(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	if response.AutoSnapshotPolicies.AutoSnapshotPolicy != nil && len(response.AutoSnapshotPolicies.AutoSnapshotPolicy) > 0 {
		return response.AutoSnapshotPolicies.AutoSnapshotPolicy[0], nil
	}

	return nil, nil
}

func (x *TableAlicloudEcsDiskGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsDiskGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("billing_method").ColumnType(schema.ColumnTypeString).Description("The billing method of the disk. Possible values are: PrePaid and PostPaid.").
			Extractor(column_value_extractor.StructSelector("DiskChargeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_retention_days").ColumnType(schema.ColumnTypeInt).Description("The retention period of the automatic snapshot.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RetentionDays")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device").ColumnType(schema.ColumnTypeString).Description("The device name of the disk on its associated instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance to which the disk is attached. This parameter has a value only when the value of Status is In_use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The device name of the disk on its associated instance.").
			Extractor(column_value_extractor.StructSelector("KMSKeyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_name").ColumnType(schema.ColumnTypeString).Description("The name of the automatic snapshot policy applied to the disk.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_code").ColumnType(schema.ColumnTypeString).Description("The product code in Alibaba Cloud Marketplace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the ECS disk.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsDiskARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).Description("The attachment information of the cloud disk.").
			Extractor(column_value_extractor.StructSelector("Attachments.Attachment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsDiskARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A user provided, human readable description for this resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_set_partition_number").ColumnType(schema.ColumnTypeInt).Description("The maximum number of partitions in a storage set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Specifies the current state of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_repeat_week_days").ColumnType(schema.ColumnTypeString).Description("The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RepeatWeekdays")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_status").ColumnType(schema.ColumnTypeString).Description("The status of the automatic snapshot policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Status")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_tags").ColumnType(schema.ColumnTypeJSON).Description("The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Tags.Tag")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_with_instance").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the disk is released when its associated instance is released.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_auto_snapshot").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the automatic snapshots of the disk are deleted when the disk is released.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_set_id").ColumnType(schema.ColumnTypeString).Description("The ID of the storage set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the disk was attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_auto_snapshot").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the automatic snapshot policy feature was enabled for the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_automated_snapshot_policy").ColumnType(schema.ColumnTypeBool).Description("Indicates whether an automatic snapshot policy was applied to the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).Description("The ID of the image used to create the instance. This parameter is empty unless the disk was created from an image. The value of this parameter remains unchanged throughout the lifecycle of the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).Description("The serial number of the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).Description("The zone name in which the resource is created.").
			Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("performance_level").ColumnType(schema.ColumnTypeString).Description("The performance level of the ESSD.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mount_instance_num").ColumnType(schema.ColumnTypeInt).Description("The number of instances to which the Shared Block Storage device is attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("portable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the disk is removable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_id").ColumnType(schema.ColumnTypeString).Description("The ID of the automatic snapshot policy applied to the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iops_read").ColumnType(schema.ColumnTypeInt).Description("The number of I/O reads per second.").
			Extractor(column_value_extractor.StructSelector("IOPSRead")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_enable_cross_region_copy").ColumnType(schema.ColumnTypeBool).Description("The ID of the automatic snapshot policy applied to the disk.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("EnableCrossRegionCopy")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expired_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the subscription disk expires.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mount_instances").ColumnType(schema.ColumnTypeJSON).Description("The attaching information of the disk.").
			Extractor(column_value_extractor.StructSelector("MountInstances.MountInstance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the disk was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category").ColumnType(schema.ColumnTypeString).Description("The category of the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detached_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the disk was detached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_creation_time").ColumnType(schema.ColumnTypeString).Description("The time when the auto snapshot policy was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("CreationTime")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iops").ColumnType(schema.ColumnTypeInt).Description("The number of input/output operations per second (IOPS).").
			Extractor(column_value_extractor.StructSelector("IOPS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot_id").ColumnType(schema.ColumnTypeString).Description("The ID of the snapshot used to create the disk. This parameter is empty unless the disk was created from a snapshot. The value of this parameter remains unchanged throughout the lifecycle of the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("Specifies the size of the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encrypted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the disk was encrypted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iops_write").ColumnType(schema.ColumnTypeInt).Description("The number of I/O writes per second.").
			Extractor(column_value_extractor.StructSelector("IOPSWrite")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the disk belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_lock").ColumnType(schema.ColumnTypeJSON).Description("The reasons why the disk was locked.").
			Extractor(column_value_extractor.StructSelector("OperationLocks.OperationLock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A friendly name for the resource.").
			Extractor(column_value_extractor.StructSelector("DiskName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Specifies the type of the disk. Possible values are: 'system' and 'data'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_snapshot_policy_time_points").ColumnType(schema.ColumnTypeString).Description("The points in time at which automatic snapshots are created. The least interval at which snapshots can be created is one hour. Valid values: 0 to 23, which corresponds to the hours of the day from 00:00 to 23:00. 1 indicates 01:00. You can specify multiple points in time.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsDiskAutoSnapshotPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TimePoints")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := ecsDiskTitle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
	}
}

func (x *TableAlicloudEcsDiskGenerator) GetSubTables() []*schema.Table {
	return nil
}
