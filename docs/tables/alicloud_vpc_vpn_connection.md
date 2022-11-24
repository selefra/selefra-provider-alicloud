# Table: alicloud_vpc_vpn_connection

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| create_time | timestamp | X | √ | The time when the IPsec-VPN connection was created. | 
| enable_nat_traversal | bool | X | √ | Indicates whether to enable the NAT traversal feature. | 
| ipsec_config | json | X | √ | The configurations for Phase 2 negotiations. | 
| vco_health_check | json | X | √ | The health check configurations. | 
| vpn_bgp_config | json | X | √ | BGP configuration information. | 
| vpn_connection_id | string | X | √ | The ID of the IPsec-VPN connection. | 
| ike_config | json | X | √ | The configurations of Phase 1 negotiations. | 
| name | string | X | √ | The name of the IPsec-VPN connection. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| remote_subnet | cidr | X | √ | The CIDR block of the on-premises data center. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| status | string | X | √ | The status of the IPsec-VPN connection. | 
| customer_gateway_id | string | X | √ | The ID of the customer gateway. | 
| effect_immediately | bool | X | √ | Indicates whether IPsec-VPN negotiations are initiated immediately. | 
| enable_dpd | bool | X | √ | Indicates whether dead peer detection (DPD) is enabled. | 
| local_subnet | cidr | X | √ | The CIDR block of the virtual private cloud (VPC). | 


