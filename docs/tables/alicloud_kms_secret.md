# Table: alicloud_kms_secret

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| secret_type | string | X | √ | The type of the secret. | 
| description | string | X | √ | The description of the secret. | 
| last_rotation_date | timestamp | X | √ | Date of last rotation of Secret. | 
| planned_delete_time | timestamp | X | √ | The time when the KMS Secret is planned to delete. | 
| extended_config | json | X | √ | The extended configuration of Secret. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the secret. | 
| create_time | timestamp | X | √ | The time when the KMS Secret was created. | 
| next_rotation_date | timestamp | X | √ | The date of next rotation of Secret. | 
| update_time | timestamp | X | √ | The time when the KMS Secret was modifies. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN). | 
| version_ids | json | X | √ | The list of secret versions. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| automatic_rotation | string | X | √ | Specifies whether automatic key rotation is enabled. | 
| encryption_key_id | string | X | √ | The ID of the KMS customer master key (CMK) that is used to encrypt the secret value. | 
| rotation_interval | string | X | √ | The rotation perion of Secret. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


