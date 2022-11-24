# Table: alicloud_ecs_security_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ | The type of the security group. Possible values are: normal, and enterprise. | 
| service_id | int | X | √ | The ID of the distributor to which the security group belongs. | 
| region | string | X | √ | The name of the region where the resource belongs. | 
| security_group_id | string | X | √ | The ID of the security group. | 
| vpc_id | string | X | √ | he ID of the VPC to which the security group belongs. | 
| creation_time | timestamp | X | √ | The time when the security group was created. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the security group belongs. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The name of the security group. | 
| service_managed | bool | X | √ | Indicates whether the user is an Alibaba Cloud service or a distributor. | 
| permissions | json | X | √ | Details about the security group rules. | 
| tags_src | json | X | √ | A list of tags attached with the security group. | 
| account_id | string | X | √ | The alicloud Account ID in which the resource is located. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS security group. | 
| description | string | X | √ | The description of the security group. | 
| inner_access_policy | string | X | √ | The description of the security group. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 


