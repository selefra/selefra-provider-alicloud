# Table: alicloud_vpc_route_table

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| creation_time | timestamp | X | √ | The time when the Route Table was created. | 
| route_table_type | string | X | √ | The type of Route Table. | 
| status | string | X | √ | The status of the route table. | 
| route_entries | json | X | √ | Route entry represents a route item of one VPC route table. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| description | string | X | √ | The description of the Route Table. | 
| vswitch_ids | json | X | √ | The unique ID of the VPC. | 
| vpc_id | string | X | √ | The ID of the VPC to which the route table belongs. | 
| tags_src | json | X | √ | A list of tags assigned to the resource. | 
| route_table_id | string | X | √ | The id of the Route Table. | 
| router_type | string | X | √ | The type of the VRouter to which the route table belongs. Valid Values are 'VRouter' and 'VBR'. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the VPC belongs. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the Route Table. | 
| router_id | string | X | √ | The ID of the region to which the VPC belongs. | 
| owner_id | string | X | √ | The ID of the owner of the VPC. | 
| tags | json | X | √ | A map of tags for the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


