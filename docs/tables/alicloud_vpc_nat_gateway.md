# Table: alicloud_vpc_nat_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_pay | bool | X | √ | Indicates whether auto pay is enabled. | 
| business_status | string | X | √ | The status of the NAT gateway. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the NAT gateway. | 
| nat_gateway_id | string | X | √ | The ID of the NAT gateway. | 
| vpc_id | string | X | √ | The ID of the virtual private cloud (VPC) to which the NAT gateway belongs. | 
| nat_gateway_private_info | json | X | √ | The information of the virtual private cloud (VPC) to which the enhanced NAT gateway belongs. | 
| snat_table_ids | json | X | √ | The ID of the SNAT table for the NAT gateway. | 
| title | string | X | √ | Title of the resource. | 
| status | string | X | √ | The state of the NAT gateway. | 
| description | string | X | √ | The description of the NAT gateway. | 
| ecs_metric_enabled | bool | X | √ | Indicates whether the traffic monitoring feature is enabled. | 
| expired_ime | timestamp | X | √ | The time when the NAT gateway expires. | 
| internet_charge_type | string | X | √ | The billing method of the NAT gateway. | 
| spec | string | X | √ | The size of the NAT gateway. | 
| forward_table_ids | json | X | √ | The ID of the Destination Network Address Translation (DNAT) table. | 
| billing_method | string | X | √ | The billing method of the NAT gateway. | 
| creation_time | timestamp | X | √ | The time when the NAT gateway was created. | 
| resource_group_id | string | X | √ | The ID of the resource group. | 
| ip_lists | json | X | √ | The elastic IP address (EIP) that is associated with the NAT gateway. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| nat_type | string | X | √ | The type of the NAT gateway. Valid values: 'Normal' and 'Enhanced'. | 
| deletion_protection | bool | X | √ | Indicates whether deletion protection is enabled. | 


