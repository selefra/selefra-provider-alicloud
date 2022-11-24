package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudEcsAutoscalingGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsAutoscalingGroupGenerator{}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetTableName() string {
	return "alicloud_ecs_autoscaling_group"
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.AutoscalingService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ess.CreateDescribeScalingGroupsRequest()
			request.Scheme = "https"
			request.PageSize = requests.NewInteger(50)
			request.PageNumber = requests.NewInteger(1)

			count := 0
			for {
				response, err := client.DescribeScalingGroups(request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, group := range response.ScalingGroups.ScalingGroup {
					resultChannel <- group
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

func getEcsAutoscalingGroupAka(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ess.ScalingGroup)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ess:" + data.RegionId + ":" + accountID + ":scalinggroup/" + data.ScalingGroupId}

	return akas, nil
}
func getEcsAutoscalingGroupConfigurations(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ess.ScalingGroup)

	client, err := alicloud_client.AutoscalingService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ess.CreateDescribeScalingConfigurationsRequest()
	request.Scheme = "https"
	request.ScalingGroupId = data.ScalingGroupId

	response, err := client.DescribeScalingConfigurations(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	if response.ScalingConfigurations.ScalingConfiguration != nil {
		return response.ScalingConfigurations.ScalingConfiguration, nil
	}

	return nil, nil
}
func getEcsAutoscalingGroupScalingInstances(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ess.ScalingGroup)

	client, err := alicloud_client.AutoscalingService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ess.CreateDescribeScalingInstancesRequest()
	request.Scheme = "https"
	request.ScalingGroupId = data.ScalingGroupId

	response, err := client.DescribeScalingInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	if response.ScalingInstances.ScalingInstance != nil {
		return response.ScalingInstances.ScalingInstance, nil
	}

	return nil, nil
}
func getEcsAutoscalingGroupTags(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ess.ScalingGroup)

	client, err := alicloud_client.AutoscalingService(ctx, clientMeta, taskClient, task)
	if err != nil {

		return nil, err
	}

	request := ess.CreateListTagResourcesRequest()
	request.Scheme = "https"
	request.ResourceType = "scalingGroup"
	request.ResourceId = &[]string{data.ScalingGroupId}

	response, err := client.ListTagResources(request)
	if serverErr, ok := err.(*errors.ServerError); ok {

		return nil, serverErr
	}

	return response, nil
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return alicloud_client.BuildRegionList()
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("on_demand_base_capacity").ColumnType(schema.ColumnTypeInt).Description("The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_type").ColumnType(schema.ColumnTypeString).Description("The health check mode of the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template_id").ColumnType(schema.ColumnTypeString).Description("The ID of the launch template used by the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("desired_capacity").ColumnType(schema.ColumnTypeInt).Description("The expected number of ECS instances in the scaling group. Auto Scaling automatically keeps the ECS instances at this number.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_size").ColumnType(schema.ColumnTypeInt).Description("The maximum number of ECS instances in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsAutoscalingGroupTags(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TagResources.TagResource")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compensate_with_on_demand").ColumnType(schema.ColumnTypeBool).Description("Specifies whether to automatically create pay-as-you-go instances to meet the requirement for the number of ECS instances in the scaling group when the number of preemptible instances cannot be reached due to reasons such as cost or insufficient resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the scaling group was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_size").ColumnType(schema.ColumnTypeInt).Description("The minimum number of ECS instances in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modification_time").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the scaling group was modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_wait_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that are in the pending state to be added in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("removing_wait_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that are in the pending state to be removed from the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_instance_remedy").ColumnType(schema.ColumnTypeBool).Description("Specifies whether to supplement preemptible instances when the target capacity of preemptible instances is not fulfilled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vserver_groups").ColumnType(schema.ColumnTypeJSON).Description("Details about backend server groups.").
			Extractor(column_value_extractor.StructSelector("VServerGroups.VServerGroup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that have been added to the scaling group and are running properly.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_cooldown").ColumnType(schema.ColumnTypeInt).Description("The default cooldown period of the scaling group (in seconds).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("ScalingGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_instances").ColumnType(schema.ColumnTypeJSON).Description("A list of ECS instances in a scaling group.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsAutoscalingGroupScalingInstances(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("life_cycle_state").ColumnType(schema.ColumnTypeString).Description("The lifecycle status of the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stopped_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of instances that are in the stopped state in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_az_policy").ColumnType(schema.ColumnTypeString).Description("The ECS instance scaling policy for a multi-zone scaling group.").
			Extractor(column_value_extractor.StructSelector("MultiAZPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("on_demand_percentage_above_base_capacity").ColumnType(schema.ColumnTypeInt).Description("The percentage of pay-as-you-go instances to be created when instances are added to the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_capacity").ColumnType(schema.ColumnTypeInt).Description("The total number of ECS instances in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_policy").ColumnType(schema.ColumnTypeString).Description("Specifies the reclaim policy of the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VPC to which the scaling group belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protected_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that are in the protected state in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("removing_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that are being removed from the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spot_instance_pools").ColumnType(schema.ColumnTypeInt).Description("The number of available instance types. Auto Scaling will create preemptible instances of multiple instance types available at the lowest cost. Valid values: 0 to 10.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("standby_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of instances that are in the standby state in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_ids").ColumnType(schema.ColumnTypeJSON).Description("The IDs of the SLB instances that are associated with the scaling group.").
			Extractor(column_value_extractor.StructSelector("LoadBalancerIds.LoadBalancerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("removal_policies").ColumnType(schema.ColumnTypeJSON).Description("Details about policies for removing ECS instances from the scaling group.").
			Extractor(column_value_extractor.StructSelector("RemovalPolicies.RemovalPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_deletion_protection").ColumnType(schema.ColumnTypeBool).Description("Indicates whether scaling group deletion protection is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_id").ColumnType(schema.ColumnTypeString).Description("The ID of the VSwitch that is associated with the scaling group.").
			Extractor(column_value_extractor.StructSelector("VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("A list of tags attached with the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getEcsAutoscalingGroupTags(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("TagResources.TagResource")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_scaling_configuration_id").ColumnType(schema.ColumnTypeString).Description("The ID of the active scaling configuration in the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template_version").ColumnType(schema.ColumnTypeString).Description("The version of the launch template used by the scaling group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_instance_ids").ColumnType(schema.ColumnTypeJSON).Description("The IDs of the ApsaraDB RDS instances that are associated with the scaling group.").
			Extractor(column_value_extractor.StructSelector("DBInstanceIds.DBInstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vswitch_ids").ColumnType(schema.ColumnTypeJSON).Description("A collection of IDs of the VSwitches that are associated with the scaling group.").
			Extractor(column_value_extractor.StructSelector("VSwitchIds.VSwitchId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_configurations").ColumnType(schema.ColumnTypeJSON).Description("A list of scaling configurations.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsAutoscalingGroupConfigurations(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getEcsAutoscalingGroupAka(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A friendly name for the resource.").
			Extractor(column_value_extractor.StructSelector("ScalingGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_group_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_capacity").ColumnType(schema.ColumnTypeInt).Description("The number of ECS instances that are being added to the scaling group, but are still being configured.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspended_processes").ColumnType(schema.ColumnTypeJSON).Description("The scaling activity that is suspended. If no scaling activity is suspended, the returned value is null.").
			Extractor(column_value_extractor.StructSelector("SuspendedProcesses.SuspendedProcess")).Build(),
	}
}

func (x *TableAlicloudEcsAutoscalingGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
