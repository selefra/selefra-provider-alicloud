# Table: alicloud_ram_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_date | timestamp | X | √ | The time when the RAM user group was created. | 
| update_date | timestamp | X | √ | The time when the RAM user group was modified. | 
| users | json | X | √ | A list of users in the group. | 
| title | string | X | √ | Title of the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the RAM user group. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the RAM user group. | 
| comments | string | X | √ | The description of the RAM user group. | 
| attached_policy | json | X | √ | A list of policies attached to a RAM user group. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


