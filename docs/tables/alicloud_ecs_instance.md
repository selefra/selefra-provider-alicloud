# Table: alicloud_ecs_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_release_time | timestamp | X | √ | The automatic release time of the pay-as-you-go instance. | 
| local_storage_amount | int | X | √ | The number of local disks attached to the instance. | 
| registration_time | timestamp | X | √ | The time when the instance is registered. | 
| serial_number | string | X | √ | The serial number of the instance. | 
| spot_price_limit | float | X | √ | The maximum hourly price for the instance. | 
| vpc_attributes | json | X | √ | The VPC attributes of the instance. | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| title | string | X | √ | Title of the resource. | 
| family | string | X | √ | The instance family of the instance. | 
| invocation_count | int | X | √ | The count of instance invocation | 
| public_ip_address | json | X | √ | The public IP addresses of instances. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS instance. | 
| cpu_options_core_count | int | X | √ | The number of CPU cores. | 
| cpu_options_numa | string | X | √ | The number of threads allocated. | 
| last_invoked_time | timestamp | X | √ | The time when the instance is last invoked. | 
| sale_cycle | string | X | √ | The billing cycle of the instance. | 
| start_time | timestamp | X | √ | The start time of the bidding mode for the preemptible instance. | 
| os_name | string | X | √ | The name of the operating system for the instance. | 
| deletion_protection | bool | X | √ | Indicates whether you can use the ECS console or call the DeleteInstance operation to release the instance. | 
| activation_id | string | X | √ | The activation Id if the instance. | 
| device_available | bool | X | √ | Indicates whether data disks can be attached to the instance. | 
| tags | json | X | √ | A map of tags for the resource. | 
| internet_max_bandwidth_in | int | X | √ | The maximum inbound bandwidth from the Internet (in Mbit/s). | 
| internet_max_bandwidth_out | int | X | √ | The maximum outbound bandwidth to the Internet (in Mbit/s). | 
| local_storage_capacity | int | X | √ | The capacity of local disks attached to the instance. | 
| inner_ip_address | json | X | √ | The internal IP addresses of classic network-type instances. This parameter takes effect when InstanceNetworkType is set to classic. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,). | 
| metadata_options | json | X | √ | The collection of metadata options. | 
| os_type | string | X | √ | The type of the operating system. Possible values are: windows and linux. | 
| eip_address | json | X | √ | The information of the EIP associated with the instance. | 
| instance_type | string | X | √ | The type of the instance. | 
| creation_time | timestamp | X | √ | The time when the instance was created. | 
| agent_version | string | X | √ | The agent version. | 
| dedicated_instance_affinity | string | X | √ | Indicates whether the instance on a dedicated host is associated with the dedicated host. | 
| ecs_capacity_reservation_preference | string | X | √ | The preference of the ECS capacity reservation. | 
| gpu_amount | int | X | √ | The number of GPUs for the instance type. | 
| io_optimized | bool | X | √ | Specifies whether the instance is I/O optimized. | 
| memory | int | X | √ | The memory size of the instance (in MiB). | 
| os_version | string | X | √ | The version of the operating system. | 
| spot_duration | int | X | √ | The protection period of the preemptible instance (in hours). | 
| vlan_id | string | X | √ | The VLAN ID of the instance. | 
| operation_locks | json | X | √ | Details about the reasons why the instance was locked. | 
| description | string | X | √ | The description of the instance. | 
| deployment_set_id | string | X | √ | The ID of the deployment set. | 
| host_name | string | X | √ | The hostname of the instance. | 
| internet_charge_type | string | X | √ | The billing method for network usage. Valid values:PayByBandwidth,PayByTraffic | 
| stopped_mode | string | X | √ | Indicates whether the instance continues to be billed after it is stopped. | 
| rdma_ip_address | json | X | √ | The RDMA IP address of HPC instance. | 
| status | string | X | √ | The status of the instance. Possible values are: Pending, Running, Starting, Stopping, and Stopped | 
| network_type | string | X | √ | The type of the network. | 
| network_interfaces | json | X | √ | Details about the ENIs bound to the instance. | 
| instance_network_type | string | X | √ | The network type of the instance. | 
| dedicated_host_cluster_id | string | X | √ | The cluster ID of the dedicated host. | 
| gpu_spec | string | X | √ | The category of GPUs for the instance type. | 
| recyclable | bool | X | √ | Indicates whether the instance can be recycled. | 
| private_ip_address | json | X | √ | The private IP addresses of instances. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| instance_id | string | X | √ | The ID of the instance. | 
| cpu | int | X | √ | The number of vCPUs. | 
| dedicated_instance_tenancy | string | X | √ | Indicates whether the instance is hosted on a dedicated host. | 
| deployment_set_group_no | int | X | √ | The group No. of the instance in a deployment set when the deployment set is used to distribute instances across multiple physical machines. | 
| ecs_capacity_reservation_id | string | X | √ | The ID of the capacity reservation. | 
| credit_specification | string | X | √ | The performance mode of the burstable instance. | 
| dedicated_host_id | string | X | √ | The ID of the dedicated host. | 
| dedicated_host_name | string | X | √ | The name of the dedicated host. | 
| key_pair_name | string | X | √ | The name of the SSH key pair for the instance. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the instance belongs. | 
| vpc_id | string | X | √ | The type of the instance. | 
| billing_method | string | X | √ | The billing method for network usage. | 
| image_id | string | X | √ | The ID of the image that the instance is running. | 
| os_name_en | string | X | √ | The English name of the operating system for the instance. | 
| spot_strategy | string | X | √ | The preemption policy for the pay-as-you-go instance. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| name | string | X | √ | The name of the instance. | 
| expired_time | timestamp | X | √ | The expiration time of the instance. | 
| hpc_cluster_id | string | X | √ | The ID of the HPC cluster to which the instance belongs. | 
| is_spot | bool | X | √ | Indicates whether the instance is a spot instance, or not. | 
| connected | bool | X | √ | Indicates whether the instance is connected.. | 
| cpu_options_threads_per_core | int | X | √ | The number of threads per core. | 
| security_group_ids | json | X | √ | The IDs of security groups to which the instance belongs. | 
| zone | string | X | √ | The zone in which the instance resides. | 
| selefra_id | string | X | √ | primary keys value md5 | 


