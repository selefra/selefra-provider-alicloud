# Table: alicloud_ecs_autoscaling_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_instance_ids | json | X | √ | The IDs of the ApsaraDB RDS instances that are associated with the scaling group. | 
| removal_policies | json | X | √ | Details about policies for removing ECS instances from the scaling group. | 
| scaling_configurations | json | X | √ | A list of scaling configurations. | 
| scaling_group_id | string | X | √ | An unique identifier for the resource. | 
| vswitch_id | string | X | √ | The ID of the VSwitch that is associated with the scaling group. | 
| active_scaling_configuration_id | string | X | √ | The ID of the active scaling configuration in the scaling group. | 
| min_size | int | X | √ | The minimum number of ECS instances in the scaling group. | 
| compensate_with_on_demand | bool | X | √ | Specifies whether to automatically create pay-as-you-go instances to meet the requirement for the number of ECS instances in the scaling group when the number of preemptible instances cannot be reached due to reasons such as cost or insufficient resources. | 
| multi_az_policy | string | X | √ | The ECS instance scaling policy for a multi-zone scaling group. | 
| removing_wait_capacity | int | X | √ | The number of ECS instances that are in the pending state to be removed from the scaling group. | 
| load_balancer_ids | json | X | √ | The IDs of the SLB instances that are associated with the scaling group. | 
| pending_capacity | int | X | √ | The number of ECS instances that are being added to the scaling group, but are still being configured. | 
| spot_instance_remedy | bool | X | √ | Specifies whether to supplement preemptible instances when the target capacity of preemptible instances is not fulfilled. | 
| total_capacity | int | X | √ | The total number of ECS instances in the scaling group. | 
| removing_capacity | int | X | √ | The number of ECS instances that are being removed from the scaling group. | 
| stopped_capacity | int | X | √ | The number of instances that are in the stopped state in the scaling group. | 
| scaling_instances | json | X | √ | A list of ECS instances in a scaling group. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | A friendly name for the resource. | 
| group_deletion_protection | bool | X | √ | Indicates whether scaling group deletion protection is enabled. | 
| launch_template_version | string | X | √ | The version of the launch template used by the scaling group. | 
| protected_capacity | int | X | √ | The number of ECS instances that are in the protected state in the scaling group. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| standby_capacity | int | X | √ | The number of instances that are in the standby state in the scaling group. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| creation_time | timestamp | X | √ | The time when the scaling group was created. | 
| default_cooldown | int | X | √ | The default cooldown period of the scaling group (in seconds). | 
| desired_capacity | int | X | √ | The expected number of ECS instances in the scaling group. Auto Scaling automatically keeps the ECS instances at this number. | 
| life_cycle_state | string | X | √ | The lifecycle status of the scaling group. | 
| vpc_id | string | X | √ | The ID of the VPC to which the scaling group belongs. | 
| suspended_processes | json | X | √ | The scaling activity that is suspended. If no scaling activity is suspended, the returned value is null. | 
| tags | json | X | √ | A map of tags for the resource. | 
| spot_instance_pools | int | X | √ | The number of available instance types. Auto Scaling will create preemptible instances of multiple instance types available at the lowest cost. Valid values: 0 to 10. | 
| vserver_groups | json | X | √ | Details about backend server groups. | 
| active_capacity | int | X | √ | The number of ECS instances that have been added to the scaling group and are running properly. | 
| health_check_type | string | X | √ | The health check mode of the scaling group. | 
| launch_template_id | string | X | √ | The ID of the launch template used by the scaling group. | 
| modification_time | timestamp | X | √ | The time when the scaling group was modified. | 
| pending_wait_capacity | int | X | √ | The number of ECS instances that are in the pending state to be added in the scaling group. | 
| vswitch_ids | json | X | √ | A collection of IDs of the VSwitches that are associated with the scaling group. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| scaling_policy | string | X | √ | Specifies the reclaim policy of the scaling group. | 
| max_size | int | X | √ | The maximum number of ECS instances in the scaling group. | 
| on_demand_base_capacity | int | X | √ | The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000. | 
| on_demand_percentage_above_base_capacity | int | X | √ | The percentage of pay-as-you-go instances to be created when instances are added to the scaling group. | 


