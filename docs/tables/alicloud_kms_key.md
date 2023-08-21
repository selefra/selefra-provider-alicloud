# Table: alicloud_kms_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the CMK. | 
| automatic_rotation | string | X | √ | Indicates whether automatic key rotation is enabled. | 
| deletion_protection | string | X | √ | Indicates whether deletion protection is enabled. | 
| key_aliases | json | X | √ | A list of aliases bound to a CMK. | 
| key_id | string | X | √ | The globally unique ID of the CMK. | 
| creator | string | X | √ | The creator of the CMK. | 
| primary_key_version | string | X | √ | The ID of the current primary key version of the symmetric CMK. | 
| tags_src | json | X | √ | A list of tags assigned to the key. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| key_spec | string | X | √ | The type of the CMK. | 
| last_rotation_date | timestamp | X | √ | The date and time the last rotation was performed. | 
| delete_date | timestamp | X | √ | The date and time the CMK is scheduled for deletion. | 
| protection_level | string | X | √ | The protection level of the CMK. | 
| material_expire_time | timestamp | X | √ | The time and date the key material for the CMK expires. | 
| origin | string | X | √ | The source of the key material for the CMK. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| key_state | string | X | √ | The status of the CMK. | 
| creation_date | timestamp | X | √ | The date and time the CMK was created. | 
| description | string | X | √ | The description of the CMK. | 
| key_usage | string | X | √ | The purpose of the CMK. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


