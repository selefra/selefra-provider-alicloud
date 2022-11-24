# Table: alicloud_kms_secret

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| secret_type | string | X | √ | The type of the secret. | 
| create_time | timestamp | X | √ | The time when the KMS Secret was created. | 
| extended_config | json | X | √ | The extended configuration of Secret. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| name | string | X | √ | The name of the secret. | 
| planned_delete_time | timestamp | X | √ | The time when the KMS Secret is planned to delete. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| description | string | X | √ | The description of the secret. | 
| update_time | timestamp | X | √ | The time when the KMS Secret was modifies. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| automatic_rotation | string | X | √ | Specifies whether automatic key rotation is enabled. | 
| encryption_key_id | string | X | √ | The ID of the KMS customer master key (CMK) that is used to encrypt the secret value. | 
| last_rotation_date | timestamp | X | √ | Date of last rotation of Secret. | 
| next_rotation_date | timestamp | X | √ | The date of next rotation of Secret. | 
| rotation_interval | string | X | √ | The rotation perion of Secret. | 
| version_ids | json | X | √ | The list of secret versions. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN). | 


