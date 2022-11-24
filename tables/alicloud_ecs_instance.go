package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudEcsInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsInstanceGenerator{}

func (x *TableAlicloudEcsInstanceGenerator) GetTableName() string {
	return "alicloud_ecs_instance"
}

func (x *TableAlicloudEcsInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeInstancesRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(100)
			request.PageNumber = requests.NewInteger(1)
			request.RegionId = taskClient.(*alicloud_client.AliCloudClient).Region

			count := 0
			for {

				response, err := client.DescribeInstances(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, instance := range response.Instances.Instance {
					resultChannel <- instance

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

func getEcsInstanceARN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	instance := result.(ecs.Instance)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + instance.RegionId + ":" + accountID + ":instance/" + instance.InstanceId

	return arn, nil
}

func (x *TableAlicloudEcsInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("auto_release_time").ColumnType(schema.ColumnTypeTimestamp).Description("The automatic release time of the pay-as-you-go instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("local_storage_amount").ColumnType(schema.ColumnTypeInt).Description("The number of local disks attached to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registration_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the instance is registered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).Description("The serial number of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_price_limit").ColumnType(schema.ColumnTypeFloat).Description("The maximum hourly price for the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_attributes").ColumnType(schema.ColumnTypeJSON).Description("The VPC attributes of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("InstanceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("family").ColumnType(schema.ColumnTypeString).Description("The instance family of the instance.").
			Extractor(column_value_extractor.StructSelector("InstanceTypeFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invocation_count").ColumnType(schema.ColumnTypeInt).Description("The count of instance invocation").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_address").ColumnType(schema.ColumnTypeJSON).Description("The public IP addresses of instances.").
			Extractor(column_value_extractor.StructSelector("PublicIpAddress.IpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsInstanceARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Description("The Alibaba Cloud Resource Name (ARN) of the ECS instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsInstanceARN(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_options_core_count").ColumnType(schema.ColumnTypeInt).Description("The number of CPU cores.").
			Extractor(column_value_extractor.StructSelector("CpuOptions.CoreCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_options_numa").ColumnType(schema.ColumnTypeString).Description("The number of threads allocated.").
			Extractor(column_value_extractor.StructSelector("CpuOptions.Numa")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_invoked_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the instance is last invoked.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sale_cycle").ColumnType(schema.ColumnTypeString).Description("The billing cycle of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Description("The start time of the bidding mode for the preemptible instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_name").ColumnType(schema.ColumnTypeString).Description("The name of the operating system for the instance.").
			Extractor(column_value_extractor.StructSelector("OSName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).Description("Indicates whether you can use the ECS console or call the DeleteInstance operation to release the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activation_id").ColumnType(schema.ColumnTypeString).Description("The activation Id if the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device_available").ColumnType(schema.ColumnTypeBool).Description("Indicates whether data disks can be attached to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.StructSelector("Tags.Tag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_max_bandwidth_in").ColumnType(schema.ColumnTypeInt).Description("The maximum inbound bandwidth from the Internet (in Mbit/s).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_max_bandwidth_out").ColumnType(schema.ColumnTypeInt).Description("The maximum outbound bandwidth to the Internet (in Mbit/s).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("local_storage_capacity").ColumnType(schema.ColumnTypeInt).Description("The capacity of local disks attached to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inner_ip_address").ColumnType(schema.ColumnTypeJSON).Description("The internal IP addresses of classic network-type instances. This parameter takes effect when InstanceNetworkType is set to classic. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,).").
			Extractor(column_value_extractor.StructSelector("InnerIpAddress.IpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_options").ColumnType(schema.ColumnTypeJSON).Description("The collection of metadata options.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_type").ColumnType(schema.ColumnTypeString).Description("The type of the operating system. Possible values are: windows and linux.").
			Extractor(column_value_extractor.StructSelector("OSType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("eip_address").ColumnType(schema.ColumnTypeJSON).Description("The information of the EIP associated with the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).Description("The type of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the instance was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_version").ColumnType(schema.ColumnTypeString).Description("The agent version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_instance_affinity").ColumnType(schema.ColumnTypeString).Description("Indicates whether the instance on a dedicated host is associated with the dedicated host.").
			Extractor(column_value_extractor.StructSelector("DedicatedInstanceAttribute.Affinity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ecs_capacity_reservation_preference").ColumnType(schema.ColumnTypeString).Description("The preference of the ECS capacity reservation.").
			Extractor(column_value_extractor.StructSelector("EcsCapacityReservationAttr.CapacityReservationPreference")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gpu_amount").ColumnType(schema.ColumnTypeInt).Description("The number of GPUs for the instance type.").
			Extractor(column_value_extractor.StructSelector("GPUAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("io_optimized").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the instance is I/O optimized.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeInt).Description("The memory size of the instance (in MiB).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_version").ColumnType(schema.ColumnTypeString).Description("The version of the operating system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_duration").ColumnType(schema.ColumnTypeInt).Description("The protection period of the preemptible instance (in hours).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vlan_id").ColumnType(schema.ColumnTypeString).Description("The VLAN ID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_locks").ColumnType(schema.ColumnTypeJSON).Description("Details about the reasons why the instance was locked.").
			Extractor(column_value_extractor.StructSelector("OperationLocks.LockReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_set_id").ColumnType(schema.ColumnTypeString).Description("The ID of the deployment set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_name").ColumnType(schema.ColumnTypeString).Description("The hostname of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internet_charge_type").ColumnType(schema.ColumnTypeString).Description("The billing method for network usage. Valid values:PayByBandwidth,PayByTraffic").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopped_mode").ColumnType(schema.ColumnTypeString).Description("Indicates whether the instance continues to be billed after it is stopped.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rdma_ip_address").ColumnType(schema.ColumnTypeJSON).Description("The RDMA IP address of HPC instance.").
			Extractor(column_value_extractor.StructSelector("RdmaIpAddress.IpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the instance. Possible values are: Pending, Running, Starting, Stopping, and Stopped").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_type").ColumnType(schema.ColumnTypeString).Description("The type of the network.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interfaces").ColumnType(schema.ColumnTypeJSON).Description("Details about the ENIs bound to the instance.").
			Extractor(column_value_extractor.StructSelector("NetworkInterfaces.NetworkInterface")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_network_type").ColumnType(schema.ColumnTypeString).Description("The network type of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_host_cluster_id").ColumnType(schema.ColumnTypeString).Description("The cluster ID of the dedicated host.").
			Extractor(column_value_extractor.StructSelector("DedicatedHostAttribute.DedicatedHostClusterId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gpu_spec").ColumnType(schema.ColumnTypeString).Description("The category of GPUs for the instance type.").
			Extractor(column_value_extractor.StructSelector("GPUSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recyclable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the instance can be recycled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_address").ColumnType(schema.ColumnTypeJSON).Description("The private IP addresses of instances.").
			Extractor(column_value_extractor.StructSelector("VpcAttributes.PrivateIpAddress.IpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The ID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu").ColumnType(schema.ColumnTypeInt).Description("The number of vCPUs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_instance_tenancy").ColumnType(schema.ColumnTypeString).Description("Indicates whether the instance is hosted on a dedicated host.").
			Extractor(column_value_extractor.StructSelector("DedicatedInstanceAttribute.Tenancy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_set_group_no").ColumnType(schema.ColumnTypeInt).Description("The group No. of the instance in a deployment set when the deployment set is used to distribute instances across multiple physical machines.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ecs_capacity_reservation_id").ColumnType(schema.ColumnTypeString).Description("The ID of the capacity reservation.").
			Extractor(column_value_extractor.StructSelector("EcsCapacityReservationAttr.CapacityReservationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credit_specification").ColumnType(schema.ColumnTypeString).Description("The performance mode of the burstable instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_host_id").ColumnType(schema.ColumnTypeString).Description("The ID of the dedicated host.").
			Extractor(column_value_extractor.StructSelector("DedicatedHostAttribute.DedicatedHostId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_host_name").ColumnType(schema.ColumnTypeString).Description("The name of the dedicated host.").
			Extractor(column_value_extractor.StructSelector("DedicatedHostAttribute.DedicatedHostName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_pair_name").ColumnType(schema.ColumnTypeString).Description("The name of the SSH key pair for the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_id").ColumnType(schema.ColumnTypeString).Description("The ID of the resource group to which the instance belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The type of the instance.").
			Extractor(column_value_extractor.StructSelector("VpcAttributes.VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_method").ColumnType(schema.ColumnTypeString).Description("The billing method for network usage.").
			Extractor(column_value_extractor.StructSelector("InstanceChargeType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).Description("The ID of the image that the instance is running.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_name_en").ColumnType(schema.ColumnTypeString).Description("The English name of the operating system for the instance.").
			Extractor(column_value_extractor.StructSelector("OSNameEn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_strategy").ColumnType(schema.ColumnTypeString).Description("The preemption policy for the pay-as-you-go instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the instance.").
			Extractor(column_value_extractor.StructSelector("InstanceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expired_time").ColumnType(schema.ColumnTypeTimestamp).Description("The expiration time of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hpc_cluster_id").ColumnType(schema.ColumnTypeString).Description("The ID of the HPC cluster to which the instance belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_spot").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the instance is a spot instance, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connected").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the instance is connected..").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_options_threads_per_core").ColumnType(schema.ColumnTypeInt).Description("The number of threads per core.").
			Extractor(column_value_extractor.StructSelector("CpuOptions.ThreadsPerCore")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_group_ids").ColumnType(schema.ColumnTypeJSON).Description("The IDs of security groups to which the instance belongs.").
			Extractor(column_value_extractor.StructSelector("SecurityGroupIds.SecurityGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).Description("The zone in which the instance resides.").
			Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
	}
}

func (x *TableAlicloudEcsInstanceGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricReadIopsHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsInstanceMetricCpuUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricReadIopsDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsInstanceMetricCpuUtilizationHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricWriteIopsHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricWriteIopsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricReadIopsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlicloudEcsDiskMetricWriteIopsDailyGenerator{}),
	}
}
