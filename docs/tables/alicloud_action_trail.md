# Table: alicloud_action_trail

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_organization_trail | bool | X | √ | Indicates whether the trail was created as a multi-account trail. | 
| event_rw | string | X | √ | The read/write type of the delivered events. | 
| sls_project_arn | string | X | √ | The ARN of the Log Service project to which events are delivered. | 
| start_logging_time | timestamp | X | √ | The most recent date and time when logging was enabled for the trail. | 
| title | string | X | √ | Title of the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the trail. | 
| role_name | string | X | √ | The name of the Resource Access Management (RAM) role that ActionTrail is allowed to assume. | 
| create_time | timestamp | X | √ | The time when the trail was created. | 
| sls_write_role_arn | string | X | √ | The ARN of the RAM role assumed by ActionTrail for delivering logs to the destination Log Service project. | 
| trail_region | string | X | √ | The regions to which the trail is applied. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| home_region | string | X | √ | The home region of the trail. | 
| oss_key_prefix | string | X | √ | The prefix of log files stored in the OSS bucket. | 
| stop_logging_time | timestamp | X | √ | The most recent date and time when logging was disabled for the trail. | 
| update_time | timestamp | X | √ | The most recent time when the configuration of the trail was updated. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| status | string | X | √ | The status of the trail. | 
| oss_bucket_name | string | X | √ | The name of the OSS bucket to which events are delivered. | 
| selefra_id | string | X | √ | primary keys value md5 | 


