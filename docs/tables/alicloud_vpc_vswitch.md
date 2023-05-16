# Table: alicloud_vpc_vswitch

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The description of the VPC. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the VPC belongs. | 
| owner_id | string | X | √ | The ID of the owner of the VPC. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| vpc_id | string | X | √ | The ID of the VPC to which the VSwitch belongs. | 
| cidr_block | cidr | X | √ | The IPv4 CIDR block of the VPC. | 
| available_ip_address_count | int | X | √ | The number of available IP addresses in the VSwitch. | 
| share_type | string | X | √ |  | 
| tags_src | json | X | √ | A map of tags for the resource. | 
| status | string | X | √ | The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available. | 
| ipv6_cidr_block | cidr | X | √ | The IPv6 CIDR block of the VPC. | 
| zone_id | string | X | √ | The zone to which the VSwitch belongs. | 
| is_default | bool | X | √ | True if the VPC is the default VPC in the region. | 
| network_acl_id | string | X | √ | A list of IDs of NAT Gateways. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| name | string | X | √ | The name of the VPC. | 
| vswitch_id | string | X | √ | The unique ID of the VPC. | 
| creation_time | timestamp | X | √ | The creation time of the VPC. | 
| route_table | json | X | √ | Details of the route table. | 
| cloud_resources | json | X | √ | The list of resources in the VSwitch. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


