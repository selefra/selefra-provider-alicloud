# Table: alicloud_vpc_network_acl

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_id | int | X | √ | The ID of the owner of the resource. | 
| region_id | string | X | √ | The name of the region where the resource resides. | 
| name | string | X | √ | The name of the network ACL. | 
| vpc_id | string | X | √ | The ID of the VPC associated with the network ACL. | 
| description | string | X | √ | The description of the network ACL. | 
| creation_time | timestamp | X | √ | The time when the network ACL was created. | 
| resources | json | X | √ | A list of associated resources. | 
| network_acl_id | string | X | √ | The ID of the network ACL. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| status | string | X | √ | The status of the network ACL. | 
| ingress_acl_entries | json | X | √ | A list of inbound rules of the network ACL. | 
| egress_acl_entries | json | X | √ | A list of outbound rules of the network ACL. | 
| selefra_id | string | X | √ | primary keys value md5 | 


