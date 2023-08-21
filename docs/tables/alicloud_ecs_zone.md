# Table: alicloud_ecs_zone

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| zone_id | string | X | √ | The zone ID. | 
| available_disk_categories | json | X | √ | The supported disk categories. The data type of this parameter is List. | 
| dedicated_host_generations | json | X | √ | The generation numbers of dedicated hosts. The data type of this parameter is List. | 
| title | string | X | √ | Title of the resource. | 
| available_resource_creation | json | X | √ | The types of the resources that can be created. The data type of this parameter is List. | 
| available_volume_categories | json | X | √ | The categories of available shared storage. The data type of this parameter is List. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| local_name | string | X | √ | The name of the zone in the local language. | 
| available_dedicated_host_types | json | X | √ | The supported types of dedicated hosts. The data type of this parameter is List. | 
| available_instance_types | json | X | √ | The instance types of instances that can be created. The data type of this parameter is List. | 
| available_resources | json | X | √ | An array consisting of ResourcesInfo data. | 
| selefra_id | string | X | √ | primary keys value md5 | 
| alicloud_ecs_region_selefra_id | string | X | √ | fk to alicloud_ecs_region.selefra_id | 


