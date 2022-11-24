# Table: alicloud_vpc_network_acl

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resources | json | X | √ | A list of associated resources. | 
| name | string | X | √ | The name of the network ACL. | 
| creation_time | timestamp | X | √ | The time when the network ACL was created. | 
| title | string | X | √ | Title of the resource. | 
| vpc_id | string | X | √ | The ID of the VPC associated with the network ACL. | 
| owner_id | int | X | √ | The ID of the owner of the resource. | 
| region_id | string | X | √ | The name of the region where the resource resides. | 
| ingress_acl_entries | json | X | √ | A list of inbound rules of the network ACL. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| network_acl_id | string | X | √ | The ID of the network ACL. | 
| status | string | X | √ | The status of the network ACL. | 
| description | string | X | √ | The description of the network ACL. | 
| egress_acl_entries | json | X | √ | A list of outbound rules of the network ACL. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 


