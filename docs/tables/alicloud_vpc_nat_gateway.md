# Table: alicloud_vpc_nat_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_pay | bool | X | √ | Indicates whether auto pay is enabled. | 
| description | string | X | √ | The description of the NAT gateway. | 
| ecs_metric_enabled | bool | X | √ | Indicates whether the traffic monitoring feature is enabled. | 
| resource_group_id | string | X | √ | The ID of the resource group. | 
| vpc_id | string | X | √ | The ID of the virtual private cloud (VPC) to which the NAT gateway belongs. | 
| snat_table_ids | json | X | √ | The ID of the SNAT table for the NAT gateway. | 
| nat_type | string | X | √ | The type of the NAT gateway. Valid values: 'Normal' and 'Enhanced'. | 
| deletion_protection | bool | X | √ | Indicates whether deletion protection is enabled. | 
| title | string | X | √ | Title of the resource. | 
| business_status | string | X | √ | The status of the NAT gateway. | 
| forward_table_ids | json | X | √ | The ID of the Destination Network Address Translation (DNAT) table. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| billing_method | string | X | √ | The billing method of the NAT gateway. | 
| creation_time | timestamp | X | √ | The time when the NAT gateway was created. | 
| status | string | X | √ | The state of the NAT gateway. | 
| expired_ime | timestamp | X | √ | The time when the NAT gateway expires. | 
| internet_charge_type | string | X | √ | The billing method of the NAT gateway. | 
| spec | string | X | √ | The size of the NAT gateway. | 
| ip_lists | json | X | √ | The elastic IP address (EIP) that is associated with the NAT gateway. | 
| nat_gateway_private_info | json | X | √ | The information of the virtual private cloud (VPC) to which the enhanced NAT gateway belongs. | 
| name | string | X | √ | The name of the NAT gateway. | 
| nat_gateway_id | string | X | √ | The ID of the NAT gateway. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


