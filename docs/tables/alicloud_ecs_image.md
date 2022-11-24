# Table: alicloud_ecs_image

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| architecture | string | X | √ | The image architecture. Possible values are: 'i386', and 'x86_64'. | 
| is_subscribed | bool | X | √ | Indicates whether you have subscribed to the image that corresponds to the specified product code. | 
| platform | string | X | √ | The platform of the operating system. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| description | string | X | √ | A user-defined, human readable description for the image. | 
| image_version | string | X | √ | The version of the image. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the image belongs. | 
| share_permissions | json | X | √ | A list of groups and accounts that the image can be shared. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| image_family | string | X | √ | The name of the image family. | 
| os_type | string | X | √ | The type of the operating system. Possible values are: windows,and linux | 
| status | string | X | √ | The status of the image. | 
| creation_time | timestamp | X | √ | The time when the image was created. | 
| os_name | string | X | √ | The Chinese name of the operating system. | 
| product_code | string | X | √ | The product code of the Alibaba Cloud Marketplace image. | 
| tags | json | X | √ | A map of tags for the resource. | 
| image_id | string | X | √ | The ID of the image that the instance is running. | 
| is_support_io_optimized | bool | X | √ | Indicates whether the image can be used on I/O optimized instances. | 
| disk_device_mappings | json | X | √ | The mappings between disks and snapshots under the image. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | A friendly name of the resource. | 
| is_support_cloud_init | bool | X | √ | Indicates whether the image supports cloud-init. | 
| os_name_en | string | X | √ | The English name of the operating system. | 
| progress | string | X | √ | The image creation progress, in percent(%). | 
| tags_src | json | X | √ | A list of tags attached with the image. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| size | int | X | √ | The size of the image (in GiB). | 
| is_self_shared | bool | X | √ | Indicates whether the image has been shared to other Alibaba Cloud accounts. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS image. | 
| image_owner_alias | string | X | √ | The alias of the image owner. Possible values are: system, self, others, marketplace. | 
| is_copied | bool | X | √ | Indicates whether the image is a copy of another image. | 
| usage | string | X | √ | Indicates whether the image has been used to create ECS instances. | 


