# Table: alicloud_vpc_vswitch

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_id | string | X | √ | The ID of the owner of the VPC. | 
| tags_src | json | X | √ | A map of tags for the resource. | 
| name | string | X | √ | The name of the VPC. | 
| ipv6_cidr_block | cidr | X | √ | The IPv6 CIDR block of the VPC. | 
| zone_id | string | X | √ | The zone to which the VSwitch belongs. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the VPC belongs. | 
| network_acl_id | string | X | √ | A list of IDs of NAT Gateways. | 
| vpc_id | string | X | √ | The ID of the VPC to which the VSwitch belongs. | 
| is_default | bool | X | √ | True if the VPC is the default VPC in the region. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| route_table | json | X | √ | Details of the route table. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| status | string | X | √ | The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available. | 
| cidr_block | cidr | X | √ | The IPv4 CIDR block of the VPC. | 
| available_ip_address_count | int | X | √ | The number of available IP addresses in the VSwitch. | 
| creation_time | timestamp | X | √ | The creation time of the VPC. | 
| share_type | string | X | √ |  | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| vswitch_id | string | X | √ | The unique ID of the VPC. | 
| description | string | X | √ | The description of the VPC. | 
| cloud_resources | json | X | √ | The list of resources in the VSwitch. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 


