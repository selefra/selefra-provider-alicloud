# Table: alicloud_vpc_ssl_vpn_server

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_multi_factor_auth | bool | X | √ | Indicates whether the multi factor authenticaton is enabled. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| title | string | X | √ | Title of the resource. | 
| create_time | timestamp | X | √ | The time when the SSL-VPN server was created. | 
| internet_ip | ip | X | √ | The public IP address. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the SSL-VPN server. | 
| client_ip_pool | cidr | X | √ | The client IP address pool. | 
| connections | int | X | √ | The total number of current connections. | 
| port | int | X | √ | The port used by the SSL-VPN server. | 
| local_subnet | string | X | √ | The CIDR block of the client. | 
| max_connections | int | X | √ | The maximum number of connections. | 
| proto | string | X | √ | The protocol used by the SSL-VPN server. | 
| ssl_vpn_server_id | string | X | √ | The ID of the SSL-VPN server. | 
| vpn_gateway_id | string | X | √ | The ID of the VPN gateway. | 
| cipher | string | X | √ | The encryption algorithm. | 
| is_compressed | bool | X | √ | Indicates whether the transmitted data is compressed. | 


