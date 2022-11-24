# Table: alicloud_ecs_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags_src | json | X | √ | A list of tags attached with the resource. | 
| creation_time | timestamp | X | √ | The time when the instance was created. | 
| dedicated_host_cluster_id | string | X | √ | The cluster ID of the dedicated host. | 
| dedicated_instance_affinity | string | X | √ | Indicates whether the instance on a dedicated host is associated with the dedicated host. | 
| image_id | string | X | √ | The ID of the image that the instance is running. | 
| invocation_count | int | X | √ | The count of instance invocation | 
| os_name_en | string | X | √ | The English name of the operating system for the instance. | 
| auto_release_time | timestamp | X | √ | The automatic release time of the pay-as-you-go instance. | 
| ecs_capacity_reservation_id | string | X | √ | The ID of the capacity reservation. | 
| zone | string | X | √ | The zone in which the instance resides. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the ECS instance. | 
| deletion_protection | bool | X | √ | Indicates whether you can use the ECS console or call the DeleteInstance operation to release the instance. | 
| family | string | X | √ | The instance family of the instance. | 
| hpc_cluster_id | string | X | √ | The ID of the HPC cluster to which the instance belongs. | 
| internet_max_bandwidth_in | int | X | √ | The maximum inbound bandwidth from the Internet (in Mbit/s). | 
| public_ip_address | json | X | √ | The public IP addresses of instances. | 
| vpc_id | string | X | √ | The type of the instance. | 
| host_name | string | X | √ | The hostname of the instance. | 
| start_time | timestamp | X | √ | The start time of the bidding mode for the preemptible instance. | 
| metadata_options | json | X | √ | The collection of metadata options. | 
| private_ip_address | json | X | √ | The private IP addresses of instances. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| agent_version | string | X | √ | The agent version. | 
| key_pair_name | string | X | √ | The name of the SSH key pair for the instance. | 
| memory | int | X | √ | The memory size of the instance (in MiB). | 
| serial_number | string | X | √ | The serial number of the instance. | 
| spot_strategy | string | X | √ | The preemption policy for the pay-as-you-go instance. | 
| eip_address | json | X | √ | The information of the EIP associated with the instance. | 
| cpu | int | X | √ | The number of vCPUs. | 
| ecs_capacity_reservation_preference | string | X | √ | The preference of the ECS capacity reservation. | 
| gpu_amount | int | X | √ | The number of GPUs for the instance type. | 
| is_spot | bool | X | √ | Indicates whether the instance is a spot instance, or not. | 
| vlan_id | string | X | √ | The VLAN ID of the instance. | 
| name | string | X | √ | The name of the instance. | 
| spot_duration | int | X | √ | The protection period of the preemptible instance (in hours). | 
| vpc_attributes | json | X | √ | The VPC attributes of the instance. | 
| description | string | X | √ | The description of the instance. | 
| local_storage_capacity | int | X | √ | The capacity of local disks attached to the instance. | 
| os_version | string | X | √ | The version of the operating system. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| spot_price_limit | float | X | √ | The maximum hourly price for the instance. | 
| credit_specification | string | X | √ | The performance mode of the burstable instance. | 
| dedicated_host_name | string | X | √ | The name of the dedicated host. | 
| deployment_set_id | string | X | √ | The ID of the deployment set. | 
| gpu_spec | string | X | √ | The category of GPUs for the instance type. | 
| registration_time | timestamp | X | √ | The time when the instance is registered. | 
| sale_cycle | string | X | √ | The billing cycle of the instance. | 
| instance_type | string | X | √ | The type of the instance. | 
| instance_network_type | string | X | √ | The network type of the instance. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| device_available | bool | X | √ | Indicates whether data disks can be attached to the instance. | 
| deployment_set_group_no | int | X | √ | The group No. of the instance in a deployment set when the deployment set is used to distribute instances across multiple physical machines. | 
| expired_time | timestamp | X | √ | The expiration time of the instance. | 
| internet_charge_type | string | X | √ | The billing method for network usage. Valid values:PayByBandwidth,PayByTraffic | 
| last_invoked_time | timestamp | X | √ | The time when the instance is last invoked. | 
| inner_ip_address | json | X | √ | The internal IP addresses of classic network-type instances. This parameter takes effect when InstanceNetworkType is set to classic. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,). | 
| network_interfaces | json | X | √ | Details about the ENIs bound to the instance. | 
| billing_method | string | X | √ | The billing method for network usage. | 
| cpu_options_threads_per_core | int | X | √ | The number of threads per core. | 
| dedicated_host_id | string | X | √ | The ID of the dedicated host. | 
| local_storage_amount | int | X | √ | The number of local disks attached to the instance. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the instance belongs. | 
| stopped_mode | string | X | √ | Indicates whether the instance continues to be billed after it is stopped. | 
| activation_id | string | X | √ | The activation Id if the instance. | 
| cpu_options_core_count | int | X | √ | The number of CPU cores. | 
| dedicated_instance_tenancy | string | X | √ | Indicates whether the instance is hosted on a dedicated host. | 
| recyclable | bool | X | √ | Indicates whether the instance can be recycled. | 
| rdma_ip_address | json | X | √ | The RDMA IP address of HPC instance. | 
| security_group_ids | json | X | √ | The IDs of security groups to which the instance belongs. | 
| instance_id | string | X | √ | The ID of the instance. | 
| connected | bool | X | √ | Indicates whether the instance is connected.. | 
| internet_max_bandwidth_out | int | X | √ | The maximum outbound bandwidth to the Internet (in Mbit/s). | 
| title | string | X | √ | Title of the resource. | 
| operation_locks | json | X | √ | Details about the reasons why the instance was locked. | 
| tags | json | X | √ | A map of tags for the resource. | 
| status | string | X | √ | The status of the instance. Possible values are: Pending, Running, Starting, Stopping, and Stopped | 
| cpu_options_numa | string | X | √ | The number of threads allocated. | 
| io_optimized | bool | X | √ | Specifies whether the instance is I/O optimized. | 
| network_type | string | X | √ | The type of the network. | 
| os_name | string | X | √ | The name of the operating system for the instance. | 
| os_type | string | X | √ | The type of the operating system. Possible values are: windows and linux. | 


