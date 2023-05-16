# Table: alicloud_vpc_vpn_connection

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_nat_traversal | bool | X | √ | Indicates whether to enable the NAT traversal feature. | 
| vpn_connection_id | string | X | √ | The ID of the IPsec-VPN connection. | 
| customer_gateway_id | string | X | √ | The ID of the customer gateway. | 
| local_subnet | cidr | X | √ | The CIDR block of the virtual private cloud (VPC). | 
| remote_subnet | cidr | X | √ | The CIDR block of the on-premises data center. | 
| vco_health_check | json | X | √ | The health check configurations. | 
| vpn_bgp_config | json | X | √ | BGP configuration information. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| status | string | X | √ | The status of the IPsec-VPN connection. | 
| effect_immediately | bool | X | √ | Indicates whether IPsec-VPN negotiations are initiated immediately. | 
| enable_dpd | bool | X | √ | Indicates whether dead peer detection (DPD) is enabled. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| ipsec_config | json | X | √ | The configurations for Phase 2 negotiations. | 
| title | string | X | √ | Title of the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the IPsec-VPN connection. | 
| create_time | timestamp | X | √ | The time when the IPsec-VPN connection was created. | 
| ike_config | json | X | √ | The configurations of Phase 1 negotiations. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


