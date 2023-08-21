# Table: alicloud_vpc_vpn_customer_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the customer gateway. | 
| customer_gateway_id | string | X | √ | The ID of the customer gateway. | 
| asn | int | X | √ | Specifies the ASN of the customer gateway. | 
| description | string | X | √ | The description of the customer gateway. | 
| ip_address | ip | X | √ | The IP address of the customer gateway. | 
| create_time | timestamp | X | √ | The time when the customer gateway was created. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 


