# Table: alicloud_ecs_snapshot

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the snapshot. | 
| encrypted | bool | X | √ | Indicates whether the snapshot was encrypted. | 
| source_disk_size | string | X | √ | The capacity of the source disk (in GiB). | 
| title | string | X | √ | Title of the resource. | 
| status | string | X | √ | Specifies the current state of the resource. | 
| usage | string | X | √ | Indicates whether the snapshot has been used to create images or disks. | 
| tags | json | X | √ | A map of tags for the resource. | 
| source_disk_id | string | X | √ | The ID of the source disk. This parameter is retained even after the source disk of the snapshot is released. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The alicloud Account ID in which the resource is located. | 
| creation_time | timestamp | X | √ | The time when the snapshot was created. | 
| description | string | X | √ | A user provided, human readable description for this resource. | 
| remain_time | int | X | √ | The remaining time required to create the snapshot (in seconds). | 
| product_code | string | X | √ | The product code of the Alibaba Cloud Marketplace image. | 
| progress | string | X | √ | The progress of the snapshot creation task. Unit: percent (%). | 
| region | string | X | √ | The region ID where the resource is located. | 
| snapshot_id | string | X | √ | An unique identifier for the resource. | 
| instant_access | bool | X | √ | Indicates whether the instant access feature is enabled. | 
| kms_key_id | string | X | √ | The ID of the KMS key used by the data disk. | 
| retention_days | int | X | √ | The number of days that an automatic snapshot can be retained. | 
| source_disk_type | string | X | √ | The category of the source disk. | 
| type | string | X | √ | The type of the snapshot. Default value: all. Possible values are: auto, user, and all. | 
| last_modified_time | timestamp | X | √ | The time when the snapshot was last changed. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the snapshot belongs. | 
| name | string | X | √ | A friendly name for the resource. | 
| serial_number | string | X | √ | The serial number of the snapshot. | 
| instant_access_retention_days | int | X | √ | Indicates the retention period of the instant access feature. After the retention per iod ends, the snapshot is automatically released. | 
| selefra_id | string | X | √ | primary keys value md5 | 


