# Table: alicloud_ram_credential_report

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_key_1_active | bool | X | √ | Indicates whether the user access key is active, or not. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| additional_access_key_3_exist | bool | X | √ | Indicates whether the user have access key, or not. | 
| additional_access_key_3_active | bool | X | √ | Indicates whether the user access key is active, or not. | 
| user_last_logon | timestamp | X | √ | Specifies the time when the user last logged in to the console. | 
| access_key_2_exist | bool | X | √ | Indicates whether the user have access key, or not. | 
| access_key_2_last_used | timestamp | X | √ | Specifies the time when the access key was most recently used to sign an Alicloud API request. | 
| additional_access_key_1_last_used | timestamp | X | √ | Specifies the time when the access key was most recently used to sign an Alicloud API request. | 
| password_active | bool | X | √ | Indicates whether the password is active, or not. | 
| access_key_2_active | bool | X | √ | Indicates whether the user access key is active, or not. | 
| additional_access_key_2_last_used | timestamp | X | √ | Specifies the time when the access key was most recently used to sign an Alicloud API request. | 
| generated_time | timestamp | X | √ | Specifies the time when the credential report has been generated. | 
| mfa_active | bool | X | √ | Indicates whether multi-factor authentication (MFA) device has been enabled for the user. | 
| user_creation_time | timestamp | X | √ | Specifies the time when the user is created. | 
| additional_access_key_3_last_used | timestamp | X | √ | Specifies the time when the access key was most recently used to sign an Alicloud API request. | 
| access_key_1_last_rotated | timestamp | X | √ | Specifies the time when the access key has been rotated. | 
| access_key_1_last_used | timestamp | X | √ | Specifies the time when the access key was most recently used to sign an Alicloud API request. | 
| additional_access_key_1_exist | bool | X | √ | Indicates whether the user have access key, or not. | 
| additional_access_key_2_exist | bool | X | √ | Indicates whether the user have access key, or not. | 
| access_key_2_last_rotated | timestamp | X | √ | Specifies the time when the access key has been rotated. | 
| additional_access_key_1_active | bool | X | √ | Indicates whether the user access key is active, or not. | 
| additional_access_key_1_last_rotated | timestamp | X | √ | Specifies the time when the access key has been rotated. | 
| additional_access_key_2_active | bool | X | √ | Indicates whether the user access key is active, or not. | 
| user_name | string | X | √ | The email of the RAM user. | 
| password_exist | bool | X | √ | Indicates whether the user have any password for logging in, or not. | 
| password_next_rotation | timestamp | X | √ | Specifies the time when the password will be rotated. | 
| access_key_1_exist | bool | X | √ | Indicates whether the user have access key, or not. | 
| additional_access_key_2_last_rotated | timestamp | X | √ | Specifies the time when the access key has been rotated. | 
| password_last_changed | timestamp | X | √ | Specifies the time when the password has been updated. | 
| additional_access_key_3_last_rotated | timestamp | X | √ | Specifies the time when the access key has been rotated. | 
| selefra_id | string | X | √ | primary keys value md5 | 


