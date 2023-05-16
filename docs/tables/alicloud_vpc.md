# Table: alicloud_vpc

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The name of the VPC. | 
| vrouter_id | string | X | √ | The ID of the VRouter. | 
| associated_cens | json | X | √ | The list of Cloud Enterprise Network (CEN) instances to which the VPC is attached. No value is returned if the VPC is not attached to any CEN instance. | 
| ipv6_cidr_blocks | json | X | √ | The IPv6 CIDR blocks of the VPC. | 
| tags_src | json | X | √ | A map of tags for the resource. | 
| status | string | X | √ | The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available. | 
| network_acl_num | string | X | √ |  | 
| classic_link_enabled | bool | X | √ | True if the ClassicLink function is enabled. | 
| vswitch_ids | json | X | √ | A list of VSwitches in the VPC. | 
| cen_status | string | X | √ | Indicates whether the VPC is attached to any Cloud Enterprise Network (CEN) instance. | 
| support_advanced_feature | bool | X | √ |  | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| vpc_id | string | X | √ | The unique ID of the VPC. | 
| creation_time | timestamp | X | √ | The creation time of the VPC. | 
| description | string | X | √ | The description of the VPC. | 
| tags | json | X | √ | A map of tags for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| ipv6_cidr_block | cidr | X | √ | The IPv6 CIDR block of the VPC. | 
| user_cidrs | json | X | √ | A list of user CIDRs. | 
| dhcp_options_set_id | string | X | √ | The ID of the DHCP options set associated to vpc. | 
| dhcp_options_set_status | string | X | √ | The status of the VPC network that is associated with the DHCP options set. Valid values: InUse and Pending | 
| cloud_resources | json | X | √ | The list of resources in the VPC. | 
| secondary_cidr_blocks | json | X | √ | A list of secondary IPv4 CIDR blocks of the VPC. | 
| advanced_resource | bool | X | √ |  | 
| nat_gateway_ids | json | X | √ | A list of IDs of NAT Gateways. | 
| route_table_ids | json | X | √ | A list of IDs of route tables. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the VPC. | 
| cidr_block | cidr | X | √ | The IPv4 CIDR block of the VPC. | 
| is_default | bool | X | √ | True if the VPC is the default VPC in the region. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the VPC belongs. | 
| owner_id | string | X | √ | The ID of the owner of the VPC. | 
| title | string | X | √ | Title of the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


