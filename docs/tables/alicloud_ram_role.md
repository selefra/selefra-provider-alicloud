# Table: alicloud_ram_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| update_date | timestamp | X | √ | The time when the RAM role was modified. | 
| assume_role_policy_document | json | X | √ | The content of the policy that specifies one or more entities entrusted to assume the RAM role. | 
| attached_policy | json | X | √ | A list of policies attached to a RAM role. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| description | string | X | √ | The description of the RAM role. | 
| max_session_duration | int | X | √ | The maximum session duration of the RAM role. | 
| create_date | timestamp | X | √ | The time when the RAM role was created. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the RAM role. | 
| role_id | string | X | √ | The ID of the RAM role. | 
| name | string | X | √ | The name of the RAM role. | 
| assume_role_policy_document_std | json | X | √ | The standard content of the policy that specifies one or more entities entrusted to assume the RAM role. | 


