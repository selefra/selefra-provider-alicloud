# Table: alicloud_ecs_network_interface

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_id | string | X | √ | The ID of the instance to which the ENI is bound. | 
| zone_id | string | X | √ | The zone ID of the ENI. | 
| account_id | string | X | √ | The alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the ENI. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| service_managed | bool | X | √ | Indicates whether the user is an Alibaba Cloud service or a distributor. | 
| security_group_ids | json | X | √ | The IDs of the security groups to which the ENI belongs. | 
| ipv6_sets | json | X | √ | The IPv6 addresses assigned to the ENI. | 
| tags | json | X | √ | A map of tags for the resource. | 
| network_interface_id | string | X | √ | An unique identifier for the ENI. | 
| owner_id | string | X | √ | The ID of the account that owns the ENI. | 
| creation_time | timestamp | X | √ | The time when the ENI was created. | 
| vpc_id | string | X | √ | The ID of the VPC to which the ENI belongs. | 
| associated_public_ip_allocation_id | string | X | √ | The allocation ID of the EIP. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| status | string | X | √ | The status of the ENI. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the ENI belongs. | 
| service_id | string | X | √ | The ID of the distributor to which the ENI belongs. | 
| description | string | X | √ | The description of the ENI. | 
| private_ip_sets | json | X | √ | The private IP addresses of the ENI. | 
| private_ip_address | ip | X | √ | The private IP address of the ENI. | 
| queue_number | int | X | √ | The number of queues supported by the ENI. | 
| associated_public_ip_address | ip | X | √ | The public IP address of the instance. | 
| attachment | json | X | √ | Attachments of the ENI | 
| type | string | X | √ | The type of the ENI. Valid values: 'Primary' and 'Secondary' | 
| mac_address | string | X | √ | The MAC address of the ENI. | 
| vswitch_id | string | X | √ | The ID of the VSwitch to which the ENI is connected. | 


