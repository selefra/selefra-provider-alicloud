# Table: alicloud_ecs_key_pair

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ | The time when the key pair was created. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The name of the key pair. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the key pair belongs. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| key_pair_finger_print | string | X | √ | The fingerprint of the key pair. | 
| selefra_id | string | X | √ | primary keys value md5 | 


