# Table: alicloud_ecs_autoscaling_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| on_demand_base_capacity | int | X | √ | The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| health_check_type | string | X | √ | The health check mode of the scaling group. | 
| launch_template_id | string | X | √ | The ID of the launch template used by the scaling group. | 
| desired_capacity | int | X | √ | The expected number of ECS instances in the scaling group. Auto Scaling automatically keeps the ECS instances at this number. | 
| max_size | int | X | √ | The maximum number of ECS instances in the scaling group. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| compensate_with_on_demand | bool | X | √ | Specifies whether to automatically create pay-as-you-go instances to meet the requirement for the number of ECS instances in the scaling group when the number of preemptible instances cannot be reached due to reasons such as cost or insufficient resources. | 
| creation_time | timestamp | X | √ | The time when the scaling group was created. | 
| min_size | int | X | √ | The minimum number of ECS instances in the scaling group. | 
| modification_time | timestamp | X | √ | The time when the scaling group was modified. | 
| pending_wait_capacity | int | X | √ | The number of ECS instances that are in the pending state to be added in the scaling group. | 
| removing_wait_capacity | int | X | √ | The number of ECS instances that are in the pending state to be removed from the scaling group. | 
| spot_instance_remedy | bool | X | √ | Specifies whether to supplement preemptible instances when the target capacity of preemptible instances is not fulfilled. | 
| vserver_groups | json | X | √ | Details about backend server groups. | 
| active_capacity | int | X | √ | The number of ECS instances that have been added to the scaling group and are running properly. | 
| default_cooldown | int | X | √ | The default cooldown period of the scaling group (in seconds). | 
| title | string | X | √ | Title of the resource. | 
| scaling_instances | json | X | √ | A list of ECS instances in a scaling group. | 
| life_cycle_state | string | X | √ | The lifecycle status of the scaling group. | 
| stopped_capacity | int | X | √ | The number of instances that are in the stopped state in the scaling group. | 
| multi_az_policy | string | X | √ | The ECS instance scaling policy for a multi-zone scaling group. | 
| on_demand_percentage_above_base_capacity | int | X | √ | The percentage of pay-as-you-go instances to be created when instances are added to the scaling group. | 
| total_capacity | int | X | √ | The total number of ECS instances in the scaling group. | 
| scaling_policy | string | X | √ | Specifies the reclaim policy of the scaling group. | 
| vpc_id | string | X | √ | The ID of the VPC to which the scaling group belongs. | 
| protected_capacity | int | X | √ | The number of ECS instances that are in the protected state in the scaling group. | 
| removing_capacity | int | X | √ | The number of ECS instances that are being removed from the scaling group. | 
| spot_instance_pools | int | X | √ | The number of available instance types. Auto Scaling will create preemptible instances of multiple instance types available at the lowest cost. Valid values: 0 to 10. | 
| standby_capacity | int | X | √ | The number of instances that are in the standby state in the scaling group. | 
| load_balancer_ids | json | X | √ | The IDs of the SLB instances that are associated with the scaling group. | 
| removal_policies | json | X | √ | Details about policies for removing ECS instances from the scaling group. | 
| group_deletion_protection | bool | X | √ | Indicates whether scaling group deletion protection is enabled. | 
| vswitch_id | string | X | √ | The ID of the VSwitch that is associated with the scaling group. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| active_scaling_configuration_id | string | X | √ | The ID of the active scaling configuration in the scaling group. | 
| launch_template_version | string | X | √ | The version of the launch template used by the scaling group. | 
| db_instance_ids | json | X | √ | The IDs of the ApsaraDB RDS instances that are associated with the scaling group. | 
| vswitch_ids | json | X | √ | A collection of IDs of the VSwitches that are associated with the scaling group. | 
| scaling_configurations | json | X | √ | A list of scaling configurations. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | A friendly name for the resource. | 
| scaling_group_id | string | X | √ | An unique identifier for the resource. | 
| pending_capacity | int | X | √ | The number of ECS instances that are being added to the scaling group, but are still being configured. | 
| suspended_processes | json | X | √ | The scaling activity that is suspended. If no scaling activity is suspended, the returned value is null. | 
| selefra_id | string | X | √ | primary keys value md5 | 


