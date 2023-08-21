# Table: alicloud_ecs_launch_template

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | A friendly name for the resource. | 
| created_by | string | X | √ | Specifies the creator of the launch template. | 
| latest_version_details | json | X | √ | Describes the configuration of latest launch template version. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| launch_template_id | string | X | √ | An unique identifier for the resource. | 
| create_time | timestamp | X | √ | The time when the launch template was created. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the launch template belongs. | 
| tags | json | X | √ | A map of tags for the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| default_version_number | int | X | √ | The default version number of the launch template. | 
| latest_version_number | int | X | √ | The latest version number of the launch template. | 
| modified_time | timestamp | X | √ | The time when the launch template was modified. | 
| title | string | X | √ | Title of the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


