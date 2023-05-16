# Table: alicloud_vpc_ssl_vpn_client_cert

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ca_cert | string | X | √ | The CA certificate. | 
| client_cert | string | X | √ | The client certificate. | 
| client_config | string | X | √ | The client configuration. | 
| ssl_vpn_client_cert_id | string | X | √ | The ID of the SSL client certificate. | 
| create_time | timestamp | X | √ | The time when the SSL client certificate was created. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name of the SSL client certificate. | 
| end_time | timestamp | X | √ | The time when the SSL client certificate expires. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| ssl_vpn_server_id | string | X | √ | The ID of the SSL-VPN server. | 
| status | string | X | √ | The status of the client certificate. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| client_key | string | X | √ | The client key. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| selefra_id | string | X | √ | primary keys value md5 | 


