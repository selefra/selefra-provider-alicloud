# Table: alicloud_ram_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_date | timestamp | X | √ | Policy creation date | 
| update_date | timestamp | X | √ | Last time when policy got updated  | 
| policy_document | json | X | √ | Contains the details about the policy. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| policy_document_std | json | X | √ | Contains the policy document in a canonical form for easier searching. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| policy_name | string | X | √ | The name of the policy. | 
| policy_type | string | X | √ | The type of the policy. Valid values: System and Custom. | 
| attachment_count | int | X | √ | The number of references to the policy. | 
| default_version | string | X | √ | Deafult version of the policy | 
| description | string | X | √ | The policy description | 
| selefra_id | string | X | √ | primary keys value md5 | 


