# Table: alicloud_security_center_version

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| asset_level | int | X | √ | The purchased quota for Security Center. | 
| is_over_balance | bool | X | √ | Indicates whether the number of existing servers exceeds your quota. | 
| user_defined_alarms | int | X | √ | Indicates whether the custom alert feature is enabled. | 
| title | string | X | √ | Title of the resource. | 
| instance_id | string | X | √ | The ID of the purchased Security Center instance. | 
| is_trial_version | bool | X | √ | Indicates whether Security Center is the free trial edition. | 
| app_white_list_auth_count | int | X | √ | The quota on the servers to which you can apply your application whitelist. | 
| last_trail_end_time | timestamp | X | √ | The time when the last free trial ends. | 
| release_time | timestamp | X | √ | The time when the Security Center instance expired. | 
| sas_log | int | X | √ | Indicates whether log analysis is purchased. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| version | string | X | √ | The purchased edition of Security Center. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| web_lock_auth_count | int | X | √ | The quota on the servers that web tamper proofing protects. | 
| sas_screen | int | X | √ | Indicates whether the security dashboard is purchased. | 
| sls_capacity | int | X | √ | The purchased capacity of log storage. | 
| web_lock | int | X | √ | Indicates whether web tamper proofing is enabled. | 
| app_white_list | int | X | √ | Indicates whether the application whitelist is enabled. | 
| selefra_id | string | X | √ | primary keys value md5 | 


