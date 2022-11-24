# Table: alicloud_security_center_version

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| version | string | X | √ | The purchased edition of Security Center. | 
| is_trial_version | bool | X | √ | Indicates whether Security Center is the free trial edition. | 
| app_white_list_auth_count | int | X | √ | The quota on the servers to which you can apply your application whitelist. | 
| sas_log | int | X | √ | Indicates whether log analysis is purchased. | 
| app_white_list | int | X | √ | Indicates whether the application whitelist is enabled. | 
| release_time | timestamp | X | √ | The time when the Security Center instance expired. | 
| web_lock | int | X | √ | Indicates whether web tamper proofing is enabled. | 
| web_lock_auth_count | int | X | √ | The quota on the servers that web tamper proofing protects. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| instance_id | string | X | √ | The ID of the purchased Security Center instance. | 
| is_over_balance | bool | X | √ | Indicates whether the number of existing servers exceeds your quota. | 
| user_defined_alarms | int | X | √ | Indicates whether the custom alert feature is enabled. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| asset_level | int | X | √ | The purchased quota for Security Center. | 
| last_trail_end_time | timestamp | X | √ | The time when the last free trial ends. | 
| sas_screen | int | X | √ | Indicates whether the security dashboard is purchased. | 
| sls_capacity | int | X | √ | The purchased capacity of log storage. | 
| title | string | X | √ | Title of the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 


