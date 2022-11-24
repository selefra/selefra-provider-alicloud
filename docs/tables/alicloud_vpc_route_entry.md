# Table: alicloud_vpc_route_entry

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ | The type of the route entry. | 
| next_hop_region_id | string | X | √ | The region where the next hop instance is deployed. | 
| next_hops | json | X | √ | The information about the next hop. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name of the route entry. | 
| route_entry_id | string | X | √ | The ID of the route entry. | 
| status | string | X | √ | The status of the route entry. | 
| private_ip_address | ip | X | √ | Specifies the private ip address for the route entry. | 
| next_hop_oppsite_instance_id | string | X | √ | The ID of the instance associated with the next hop. | 
| next_hop_oppsite_type | string | X | √ | The type of the next hop. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| route_table_id | string | X | √ | The ID of the route table. | 
| description | string | X | √ | The description of the VRouter. | 
| instance_id | string | X | √ | The ID of the instance associated with the next hop. | 
| next_hop_type | string | X | √ | The type of the next hop. | 
| ip_version | string | X | √ | The version of the IP protocol. | 
| destination_cidr_block | cidr | X | √ | The destination Classless Inter-Domain Routing (CIDR) block of the route entry. | 
| next_hop_oppsite_region_id | string | X | √ | The region where the next hop instance is deployed. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 


