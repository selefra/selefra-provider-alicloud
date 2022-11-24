# Table: alicloud_vpc_dhcp_options_set

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The description for the DHCP option set. | 
| domain_name_servers | string | X | √ | The IP addresses of your DNS servers. | 
| owner_id | string | X | √ | The ID of the account to which the DHCP options set belongs. | 
| domain_name | string | X | √ | The root domain. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| name | string | X | √ | The name of the DHCP option set. | 
| dhcp_options_set_id | string | X | √ | The ID of the DHCP option set. | 
| associate_vpc_count | int | X | √ | The number of VPCs associated with DHCP option set. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| status | string | X | √ | The status of the DHCP option set. | 
| associate_vpcs | json | X | √ | The information of the VPC network that is associated with the DHCP options set. | 
| title | string | X | √ | Title of the resource. | 
| boot_file_name | string | X | √ | The boot file name of DHCP option set. | 
| tftp_server_name | string | X | √ | The tftp server name of the DHCP option set. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 


