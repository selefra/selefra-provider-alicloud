# Table: alicloud_ram_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| groups | json | X | √ | A list of groups attached to the user. | 
| title | string | X | √ | Title of the resource. | 
| user_id | string | X | √ | The unique ID of the RAM user. | 
| email | string | X | √ | The email address of the RAM user. | 
| mfa_enabled | bool | X | √ | The MFA status of the user | 
| cs_user_permissions | json | X | √ | User permissions for Container Service Kubernetes clusters. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the RAM user. | 
| attached_policy | json | X | √ | A list of policies attached to a RAM user. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The username of the RAM user. | 
| last_login_date | timestamp | X | √ | The time when the RAM user last logged on to the console by using the password. | 
| mobile_phone | string | X | √ | The mobile phone number of the RAM user. | 
| mfa_device_serial_number | string | X | √ | The serial number of the MFA device. | 
| virtual_mfa_devices | json | X | √ | The list of MFA devices. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| display_name | string | X | √ | The display name of the RAM user. | 
| comments | string | X | √ | The description of the RAM user. | 
| create_date | timestamp | X | √ | The time when the RAM user was created. | 
| update_date | timestamp | X | √ | The time when the RAM user was modified. | 


