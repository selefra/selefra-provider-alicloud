# Table: alicloud_cms_monitor_host

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| nat_ip | string | X | √ | The IP address of the Network Address Translation (NAT) gateway. | 
| operating_system | string | X | √ | The operating system. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| is_aliyun_host | bool | X | √ | Indicates whether the host is provided by Alibaba Cloud. | 
| eip_id | string | X | √ | The ID of the EIP. | 
| ip_group | string | X | √ | The IP address of the host. | 
| network_type | string | X | √ | The type of the network. | 
| serial_number | string | X | √ | The serial number of the host. A host that is not provided by Alibaba Cloud has a serial number instead of an instance ID. | 
| monitoring_agent_status | json | X | √ | The status of the Cloud Monitor agent. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| host_name | string | X | √ | The name of the host. | 
| eip_address | string | X | √ | The elastic IP address (EIP) of the host. | 
| agent_version | string | X | √ | The version of the Cloud Monitor agent. | 
| ali_uid | int | X | √ | The ID of the Alibaba Cloud account. | 
| instance_id | string | X | √ | The ID of the instance. | 
| instance_type_family | string | X | √ | The type of the ECS instance. | 


