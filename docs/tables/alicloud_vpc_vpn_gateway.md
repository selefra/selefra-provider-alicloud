# Table: alicloud_vpc_vpn_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ssl_max_connections | int | X | √ | The maximum number of concurrent SSL-VPN connections. | 
| ssl_vpn | string | X | √ | Indicates whether the SSL-VPN feature is enabled. | 
| tag | string | X | √ | The tag of the VPN gateway. | 
| vswitch_id | string | X | √ | The ID of the VSwitch to which the VPN gateway belongs. | 
| vpc_id | string | X | √ | The ID of the VPC for which the VPN gateway is created. | 
| status | string | X | √ | The status of the VPN gateway. | 
| enable_bgp | bool | X | √ | Indicates whether bgp is enabled. | 
| ipsec_vpn | string | X | √ | Indicates whether the IPsec-VPN feature is enabled. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| spec | string | X | √ | The maximum bandwidth of the VPN gateway. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| name | string | X | √ | The name of the VPN gateway. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| internet_ip | ip | X | √ | The public IP address of the VPN gateway. | 
| reservation_data | json | X | √ | A set of reservation details. | 
| business_status | string | X | √ | The business state of the VPN gateway. | 
| create_time | timestamp | X | √ | The time when the VPN gateway was created. | 
| end_time | timestamp | X | √ | The creation time of the VPC. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| auto_propagate | bool | X | √ | Indicates whether auto propagate is enabled, or not. | 
| billing_method | string | X | √ | The billing method of the VPN gateway. | 
| description | string | X | √ | The description of the VPN gateway. | 


