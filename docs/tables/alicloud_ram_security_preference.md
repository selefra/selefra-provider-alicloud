# Table: alicloud_ram_security_preference

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allow_user_to_manage_mfa_devices | bool | X | √ | Indicates whether RAM users can manage their MFA devices. | 
| enable_save_mfa_ticket | bool | X | √ | Indicates whether RAM users can save security codes for multi-factor authentication (MFA) during logon. Each security code is valid for seven days. | 
| login_network_masks | json | X | √ | The subnet mask that indicates the IP addresses from which logon to the Alibaba Cloud Management Console is allowed. This parameter applies to password-based logon and single sign-on (SSO). However, this parameter does not apply to API calls that are authenticated based on AccessKey pairs. May be more than one CIDR range. If empty then login is allowed from any source. | 
| login_session_duration | int | X | √ | The validity period of a logon session of a RAM user. Unit: hours. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| allow_user_to_change_password | bool | X | √ | Indicates whether RAM users can change their passwords. | 
| allow_user_to_manage_access_keys | bool | X | √ | Indicates whether RAM users can manage their AccessKey pairs. | 
| allow_user_to_manage_public_keys | bool | X | √ | Indicates whether RAM users can manage their public keys. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 


