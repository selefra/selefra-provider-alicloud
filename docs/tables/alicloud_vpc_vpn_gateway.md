# Table: alicloud_vpc_vpn_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| business_status | string | X | √ | The business state of the VPN gateway. | 
| create_time | timestamp | X | √ | The time when the VPN gateway was created. | 
| description | string | X | √ | The description of the VPN gateway. | 
| tag | string | X | √ | The tag of the VPN gateway. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| name | string | X | √ | The name of the VPN gateway. | 
| status | string | X | √ | The status of the VPN gateway. | 
| ipsec_vpn | string | X | √ | Indicates whether the IPsec-VPN feature is enabled. | 
| internet_ip | ip | X | √ | The public IP address of the VPN gateway. | 
| spec | string | X | √ | The maximum bandwidth of the VPN gateway. | 
| ssl_vpn | string | X | √ | Indicates whether the SSL-VPN feature is enabled. | 
| vswitch_id | string | X | √ | The ID of the VSwitch to which the VPN gateway belongs. | 
| vpc_id | string | X | √ | The ID of the VPC for which the VPN gateway is created. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| auto_propagate | bool | X | √ | Indicates whether auto propagate is enabled, or not. | 
| enable_bgp | bool | X | √ | Indicates whether bgp is enabled. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| reservation_data | json | X | √ | A set of reservation details. | 
| tags | json | X | √ | A map of tags for the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| billing_method | string | X | √ | The billing method of the VPN gateway. | 
| end_time | timestamp | X | √ | The creation time of the VPC. | 
| ssl_max_connections | int | X | √ | The maximum number of concurrent SSL-VPN connections. | 
| selefra_id | string | X | √ | primary keys value md5 | 


