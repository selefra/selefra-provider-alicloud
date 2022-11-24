# Table: alicloud_rds_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| collation | string | X | √ | The character set collation of the instance. | 
| auto_upgrade_minor_version | string | X | √ | The method that is used to update the minor engine version of the instance. | 
| vpc_cloud_instance_id | string | X | √ | The ID of the cloud instance on which the specified VPC is deployed. | 
| expire_time | timestamp | X | √ | Instance expire time | 
| guard_db_instance_id | string | X | √ | The ID of the disaster recovery instance that is attached to the instance if a disaster recovery instance is deployed. | 
| dedicated_host_group_id | string | X | √ | The ID of the dedicated cluster to which the instances belong if the instances are created in a dedicated cluster. | 
| temp_db_instance_id | string | X | √ | The ID of the temporary instance that is attached to the instance if a temporary instance is deployed. | 
| temp_upgrade_time_start | string | X | √ |  | 
| availability_value | string | X | √ | The availability status of the instance. Unit: %. | 
| db_instance_storage_type | string | X | √ | The type of storage media that is used by the instance. | 
| multiple_temp_upgrade | bool | X | √ |  | 
| latest_kernel_version | string | X | √ |  | 
| account_type | string | X | √ |  | 
| master_instance_id | string | X | √ | The ID of the primary instance to which the instance is attached. If this parameter is not returned, the instance is a primary instance. | 
| advanced_features | string | X | √ | An array that consists of advanced features. The advanced features are separated by commas (,). This parameter is supported only for instances that run SQL Server. | 
| db_max_quantity | int | X | √ | The maximum number of databases that can be created on the instance. | 
| dispense_mode | string | X | √ |  | 
| account_max_quantity | int | X | √ | The maximum number of accounts that can be created on the instance. | 
| readonly_db_instance_ids | json | X | √ | An array that consists of the IDs of the read-only instances attached to the primary instance. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the instances belong. | 
| db_instance_net_type | string | X | √ | The ID of the resource group to which the VPC belongs. | 
| increment_source_db_instance_id | string | X | √ | The ID of the instance from which incremental data comes. The incremental data of a disaster recovery or read-only instance comes from its primary instance. If this parameter is not returned, the instance is a primary instance. | 
| temp_upgrade_recovery_class | string | X | √ |  | 
| temp_upgrade_recovery_cpu | int | X | √ |  | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the RDS instance. | 
| engine | string | X | √ | The database engine that the instances run. | 
| db_instance_type | string | X | √ | The role of the instances. | 
| connection_mode | string | X | √ | The connection mode of the instances. | 
| time_zone | string | X | √ | The time zone of the instance. | 
| port | string | X | √ | The internal port of the instance. | 
| connection_string | string | X | √ | The internal endpoint of the instance. | 
| ins_id | int | X | √ |  | 
| vswitch_id | string | X | √ | The ID of the vSwitch associated with the specified VPC. | 
| db_instance_disk_used | string | X | √ |  | 
| support_create_super_account | string | X | √ |  | 
| temp_upgrade_recovery_memory | int | X | √ |  | 
| origin_configuration | string | X | √ |  | 
| security_ip_mode | string | X | √ | The network isolation mode of the instance. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| engine_version | string | X | √ | The version of the database engine that the instances run. | 
| lock_mode | string | X | √ | The lock mode of the instance. | 
| temp_upgrade_recovery_time | string | X | √ |  | 
| support_upgrade_account_type | string | X | √ |  | 
| temp_upgrade_time_end | string | X | √ |  | 
| title | string | X | √ | Title of the resource. | 
| pay_type | string | X | √ | The billing method of the instances. | 
| instance_network_type | string | X | √ | The network type of the instances. | 
| db_instance_status | string | X | √ | The status of the instances | 
| temp_upgrade_recovery_max_connections | string | X | √ |  | 
| ssl_status | string | X | √ | The SSL encryption status of the Instance | 
| security_ips | json | X | √ | An array that consists of IP addresses in the IP address whitelist. | 
| security_ips_src | json | X | √ | An array that consists of IP details. | 
| sql_collector_policy | json | X | √ | The status of the SQL Explorer (SQL Audit) feature. | 
| lock_reason | string | X | √ | The reason why the instance is locked. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| db_instance_class | string | X | √ | The instance type of the instances. | 
| db_instance_id | string | X | √ | The ID of the single instance to query. | 
| region_id | string | X | √ | The ID of the region to which the instances belong. | 
| console_version | string | X | √ | The type of proxy that is enabled on the instance. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| vpc_id | string | X | √ | The ID of the VPC to which the instances belong. | 
| temp_upgrade_recovery_max_iops | string | X | √ |  | 
| tde_status | string | X | √ | The TDE status at the instance level. Valid values: Enable | Disable. | 
| tags | json | X | √ | A map of tags for the resource. | 
| category | string | X | √ | The RDS edition of the instance. | 
| zone_id | string | X | √ | The ID of the zone to which the instances belong. | 
| max_iops | int | X | √ | The maximum number of I/O requests that the instance can process per second. | 
| ip_type | string | X | √ |  | 
| db_instance_description | string | X | √ | The description of the DB Instance. | 
| db_instance_memory | float | X | √ | The memory capacity of the instance. Unit: MB. | 
| maintain_time | string | X | √ | The maintenance window of the instance. The maintenance window is displayed in UTC+8 in the ApsaraDB RDS console. | 
| db_instance_storage | int | X | √ | The type of storage media that is used by the instance. | 
| proxy_type | int | X | √ | The type of proxy that is enabled on the instance. | 
| parameters | json | X | √ | The list of running parameters for the instance. | 
| tags_src | json | X | √ | A map of tags for the resource. | 
| sql_collector_retention | int | X | √ | The log backup retention duration that is allowed by the SQL explorer feature on the instance. | 
| db_instance_cpu | string | X | √ | The number of CPUs that are configured for the instance. | 
| max_connections | int | X | √ | The maximum number of concurrent connections that are allowed by the instance. | 
| super_permission_mode | string | X | √ | Indicates whether the instance supports superuser accounts, such as the system administrator (SA) account, Active Directory (AD) account, and host account. | 
| creation_time | timestamp | X | √ | The creation time of the Instance. | 


