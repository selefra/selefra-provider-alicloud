# Table: alicloud_ecs_image

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| os_name_en | string | X | √ | The English name of the operating system. | 
| title | string | X | √ | Title of the resource. | 
| image_family | string | X | √ | The name of the image family. | 
| is_subscribed | bool | X | √ | Indicates whether you have subscribed to the image that corresponds to the specified product code. | 
| platform | string | X | √ | The platform of the operating system. | 
| progress | string | X | √ | The image creation progress, in percent(%). | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| architecture | string | X | √ | The image architecture. Possible values are: 'i386', and 'x86_64'. | 
| is_copied | bool | X | √ | Indicates whether the image is a copy of another image. | 
| is_self_shared | bool | X | √ | Indicates whether the image has been shared to other Alibaba Cloud accounts. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| usage | string | X | √ | Indicates whether the image has been used to create ECS instances. | 
| name | string | X | √ | A friendly name of the resource. | 
| image_version | string | X | √ | The version of the image. | 
| is_support_cloud_init | bool | X | √ | Indicates whether the image supports cloud-init. | 
| os_name | string | X | √ | The Chinese name of the operating system. | 
| os_type | string | X | √ | The type of the operating system. Possible values are: windows,and linux | 
| disk_device_mappings | json | X | √ | The mappings between disks and snapshots under the image. | 
| share_permissions | json | X | √ | A list of groups and accounts that the image can be shared. | 
| image_id | string | X | √ | The ID of the image that the instance is running. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS image. | 
| image_owner_alias | string | X | √ | The alias of the image owner. Possible values are: system, self, others, marketplace. | 
| creation_time | timestamp | X | √ | The time when the image was created. | 
| product_code | string | X | √ | The product code of the Alibaba Cloud Marketplace image. | 
| tags_src | json | X | √ | A list of tags attached with the image. | 
| size | int | X | √ | The size of the image (in GiB). | 
| status | string | X | √ | The status of the image. | 
| description | string | X | √ | A user-defined, human readable description for the image. | 
| is_support_io_optimized | bool | X | √ | Indicates whether the image can be used on I/O optimized instances. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the image belongs. | 
| selefra_id | string | X | √ | primary keys value md5 | 


