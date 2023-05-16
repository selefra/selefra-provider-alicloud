# Table: alicloud_ecs_security_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_group_id | string | X | √ | The ID of the resource group to which the security group belongs. | 
| permissions | json | X | √ | Details about the security group rules. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The name of the security group. | 
| security_group_id | string | X | √ | The ID of the security group. | 
| type | string | X | √ | The type of the security group. Possible values are: normal, and enterprise. | 
| vpc_id | string | X | √ | he ID of the VPC to which the security group belongs. | 
| tags_src | json | X | √ | A list of tags attached with the security group. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS security group. | 
| creation_time | timestamp | X | √ | The time when the security group was created. | 
| service_id | big_int | X | √ | The ID of the distributor to which the security group belongs. | 
| tags | json | X | √ | A map of tags for the resource. | 
| description | string | X | √ | The description of the security group. | 
| inner_access_policy | string | X | √ | The description of the security group. | 
| service_managed | bool | X | √ | Indicates whether the user is an Alibaba Cloud service or a distributor. | 
| region | string | X | √ | The name of the region where the resource belongs. | 
| account_id | string | X | √ | The alicloud Account ID in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


