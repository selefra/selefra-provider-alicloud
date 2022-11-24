# Table: alicloud_ram_access_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| user_name | string | X | √ | Name of the User that the access key belongs to. | 
| access_key_id | string | X | √ | The AccessKey ID. | 
| status | string | X | √ | The status of the AccessKey pair. Valid values: Active and Inactive. | 
| create_date | timestamp | X | √ | The time when the AccessKey pair was created. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 


