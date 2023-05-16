# Table: alicloud_vpc_eip

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| expired_time | timestamp | X | √ | The expiration time of the EIP. | 
| bandwidth | string | X | √ | The peak bandwidth of the EIP. Unit: Mbit/s. | 
| operation_locks_reason | json | X | √ | The reason why the EIP is locked. Valid values: financial security. | 
| name | string | X | √ | The name of the EIP. | 
| allocation_time | timestamp | X | √ | The time when the EIP was created. | 
| status | string | X | √ | The status of the EIP. | 
| hd_monitor_status | string | X | √ | Indicates whether fine-grained monitoring is enabled for the EIP. | 
| has_reservation_data | bool | X | √ | Indicates whether renewal data is included. | 
| instance_type | string | X | √ | The type of the instance to which the EIP is bound. | 
| title | string | X | √ | Title of the resource. | 
| available_regions | json | X | √ | The ID of the region to which the EIP belongs. | 
| instance_region_id | string | X | √ | The region ID of the bound resource. | 
| resource_group_id | string | X | √ | The ID of the resource group. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the EIP. | 
| charge_type | string | X | √ | The billing method of the EIP | 
| segment_instance_id | string | X | √ | The ID of the instance with which the contiguous EIP is associated. | 
| description | string | X | √ | The description of the EIP. | 
| isp | string | X | √ | The Internet service provider. | 
| bandwidth_package_bandwidth | string | X | √ | The maximum bandwidth of the EIP in Mbit/s. | 
| bandwidth_package_type | string | X | √ | The bandwidth value of the EIP Bandwidth Plan to which the EIP is added. | 
| mode | string | X | √ | The type of the instance to which you want to bind the EIP. | 
| netmode | string | X | √ | The network type of the EIP. | 
| private_ip_address | bool | X | √ |  | 
| second_limited | bool | X | √ | Indicates whether level-2 traffic throttling is configured. | 
| instance_id | string | X | √ | The ID of the instance to which the EIP is bound. | 
| ip_address | ip | X | √ | The IP address of the EIP. | 
| internet_charge_type | string | X | √ | The metering method of the EIP can be one of PayByBandwidth or PayByTraffic. | 
| service_managed | int | X | √ |  | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| allocation_id | string | X | √ | The unique ID of the EIP. | 
| selefra_id | string | X | √ | primary keys value md5 | 


