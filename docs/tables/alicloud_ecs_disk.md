# Table: alicloud_ecs_disk

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| billing_method | string | X | √ | The billing method of the disk. Possible values are: PrePaid and PostPaid. | 
| auto_snapshot_policy_retention_days | int | X | √ | The retention period of the automatic snapshot. | 
| device | string | X | √ | The device name of the disk on its associated instance. | 
| instance_id | string | X | √ | The ID of the instance to which the disk is attached. This parameter has a value only when the value of Status is In_use. | 
| kms_key_id | string | X | √ | The device name of the disk on its associated instance. | 
| auto_snapshot_policy_name | string | X | √ | The name of the automatic snapshot policy applied to the disk. | 
| product_code | string | X | √ | The product code in Alibaba Cloud Marketplace. | 
| disk_id | string | X | √ | An unique identifier for the resource. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS disk. | 
| attachments | json | X | √ | The attachment information of the cloud disk. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| description | string | X | √ | A user provided, human readable description for this resource. | 
| storage_set_partition_number | int | X | √ | The maximum number of partitions in a storage set. | 
| status | string | X | √ | Specifies the current state of the resource. | 
| auto_snapshot_policy_repeat_week_days | string | X | √ | The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified. | 
| auto_snapshot_policy_status | string | X | √ | The status of the automatic snapshot policy. | 
| auto_snapshot_policy_tags | json | X | √ | The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified. | 
| delete_with_instance | bool | X | √ | Indicates whether the disk is released when its associated instance is released. | 
| delete_auto_snapshot | bool | X | √ | Indicates whether the automatic snapshots of the disk are deleted when the disk is released. | 
| storage_set_id | string | X | √ | The ID of the storage set. | 
| attached_time | timestamp | X | √ | The time when the disk was attached. | 
| enable_auto_snapshot | bool | X | √ | Indicates whether the automatic snapshot policy feature was enabled for the disk. | 
| enable_automated_snapshot_policy | bool | X | √ | Indicates whether an automatic snapshot policy was applied to the disk. | 
| image_id | string | X | √ | The ID of the image used to create the instance. This parameter is empty unless the disk was created from an image. The value of this parameter remains unchanged throughout the lifecycle of the disk. | 
| serial_number | string | X | √ | The serial number of the disk. | 
| zone | string | X | √ | The zone name in which the resource is created. | 
| performance_level | string | X | √ | The performance level of the ESSD. | 
| mount_instance_num | int | X | √ | The number of instances to which the Shared Block Storage device is attached. | 
| portable | bool | X | √ | Indicates whether the disk is removable. | 
| auto_snapshot_policy_id | string | X | √ | The ID of the automatic snapshot policy applied to the disk. | 
| iops_read | int | X | √ | The number of I/O reads per second. | 
| auto_snapshot_policy_enable_cross_region_copy | bool | X | √ | The ID of the automatic snapshot policy applied to the disk. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| expired_time | timestamp | X | √ | The time when the subscription disk expires. | 
| mount_instances | json | X | √ | The attaching information of the disk. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| creation_time | timestamp | X | √ | The time when the disk was created. | 
| category | string | X | √ | The category of the disk. | 
| detached_time | timestamp | X | √ | The time when the disk was detached. | 
| auto_snapshot_policy_creation_time | string | X | √ | The time when the auto snapshot policy was created. | 
| iops | int | X | √ | The number of input/output operations per second (IOPS). | 
| tags | json | X | √ | A map of tags for the resource. | 
| source_snapshot_id | string | X | √ | The ID of the snapshot used to create the disk. This parameter is empty unless the disk was created from a snapshot. The value of this parameter remains unchanged throughout the lifecycle of the disk. | 
| size | int | X | √ | Specifies the size of the disk. | 
| encrypted | bool | X | √ | Indicates whether the disk was encrypted. | 
| iops_write | int | X | √ | The number of I/O writes per second. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the disk belongs. | 
| operation_lock | json | X | √ | The reasons why the disk was locked. | 
| name | string | X | √ | A friendly name for the resource. | 
| type | string | X | √ | Specifies the type of the disk. Possible values are: 'system' and 'data'. | 
| auto_snapshot_policy_time_points | string | X | √ | The points in time at which automatic snapshots are created. The least interval at which snapshots can be created is one hour. Valid values: 0 to 23, which corresponds to the hours of the day from 00:00 to 23:00. 1 indicates 01:00. You can specify multiple points in time. | 
| title | string | X | √ | Title of the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


