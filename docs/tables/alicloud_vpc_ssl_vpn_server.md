# Table: alicloud_vpc_ssl_vpn_server

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_compressed | bool | X | √ | Indicates whether the transmitted data is compressed. | 
| title | string | X | √ | Title of the resource. | 
| ssl_vpn_server_id | string | X | √ | The ID of the SSL-VPN server. | 
| connections | int | X | √ | The total number of current connections. | 
| client_ip_pool | cidr | X | √ | The client IP address pool. | 
| create_time | timestamp | X | √ | The time when the SSL-VPN server was created. | 
| enable_multi_factor_auth | bool | X | √ | Indicates whether the multi factor authenticaton is enabled. | 
| internet_ip | ip | X | √ | The public IP address. | 
| proto | string | X | √ | The protocol used by the SSL-VPN server. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| cipher | string | X | √ | The encryption algorithm. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| local_subnet | string | X | √ | The CIDR block of the client. | 
| port | int | X | √ | The port used by the SSL-VPN server. | 
| name | string | X | √ | The name of the SSL-VPN server. | 
| max_connections | int | X | √ | The maximum number of connections. | 
| selefra_id | string | X | √ | primary keys value md5 | 


