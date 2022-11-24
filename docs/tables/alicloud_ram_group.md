# Table: alicloud_ram_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| name | string | X | √ | The name of the RAM user group. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the RAM user group. | 
| create_date | timestamp | X | √ | The time when the RAM user group was created. | 
| update_date | timestamp | X | √ | The time when the RAM user group was modified. | 
| attached_policy | json | X | √ | A list of policies attached to a RAM user group. | 
| users | json | X | √ | A list of users in the group. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| comments | string | X | √ | The description of the RAM user group. | 
| title | string | X | √ | Title of the resource. | 


