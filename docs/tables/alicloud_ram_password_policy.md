# Table: alicloud_ram_password_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| hard_expiry | bool | X | √ | Indicates whether the password has expired. | 
| require_lowercase_characters | bool | X | √ | Indicates whether a password must contain one or more lowercase letters. | 
| require_uppercase_characters | bool | X | √ | Indicates whether a password must contain one or more uppercase letters. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| max_login_attempts | int | X | √ | The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password. | 
| max_password_age | int | X | √ | The number of days for which a password is valid. Default value: 0. The default value indicates that the password never expires. | 
| minimum_password_length | int | X | √ | The minimum required number of characters in a password. | 
| password_reuse_prevention | int | X | √ | The number of previous passwords that the user is prevented from reusing. Default value: 0. The default value indicates that the RAM user is not prevented from reusing previous passwords. | 
| require_numbers | bool | X | √ | Indicates whether a password must contain one or more digits. | 
| require_symbols | bool | X | √ | Indicates whether a password must contain one or more special characters. | 
| selefra_id | string | X | √ | primary keys value md5 | 


