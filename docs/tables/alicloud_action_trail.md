# Table: alicloud_action_trail

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| oss_key_prefix | string | X | √ | The prefix of log files stored in the OSS bucket. | 
| trail_region | string | X | √ | The regions to which the trail is applied. | 
| is_organization_trail | bool | X | √ | Indicates whether the trail was created as a multi-account trail. | 
| event_rw | string | X | √ | The read/write type of the delivered events. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| name | string | X | √ | The name of the trail. | 
| home_region | string | X | √ | The home region of the trail. | 
| create_time | timestamp | X | √ | The time when the trail was created. | 
| start_logging_time | timestamp | X | √ | The most recent date and time when logging was enabled for the trail. | 
| update_time | timestamp | X | √ | The most recent time when the configuration of the trail was updated. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| role_name | string | X | √ | The name of the Resource Access Management (RAM) role that ActionTrail is allowed to assume. | 
| status | string | X | √ | The status of the trail. | 
| sls_write_role_arn | string | X | √ | The ARN of the RAM role assumed by ActionTrail for delivering logs to the destination Log Service project. | 
| stop_logging_time | timestamp | X | √ | The most recent date and time when logging was disabled for the trail. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| oss_bucket_name | string | X | √ | The name of the OSS bucket to which events are delivered. | 
| sls_project_arn | string | X | √ | The ARN of the Log Service project to which events are delivered. | 


