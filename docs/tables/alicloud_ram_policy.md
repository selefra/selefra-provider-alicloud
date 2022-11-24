# Table: alicloud_ram_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_document_std | json | X | √ | Contains the policy document in a canonical form for easier searching. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| attachment_count | int | X | √ | The number of references to the policy. | 
| default_version | string | X | √ | Deafult version of the policy | 
| update_date | timestamp | X | √ | Last time when policy got updated  | 
| description | string | X | √ | The policy description | 
| policy_document | json | X | √ | Contains the details about the policy. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| policy_name | string | X | √ | The name of the policy. | 
| policy_type | string | X | √ | The type of the policy. Valid values: System and Custom. | 
| create_date | timestamp | X | √ | Policy creation date | 


