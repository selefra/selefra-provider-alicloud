# Table: alicloud_cs_kubernetes_cluster

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service_discovery_types | string | X | √ |  | 
| vswitch_id | string | X | √ | The IDs of VSwitches. | 
| worker_ram_role_name | string | X | √ | The name of the RAM role for worker nodes in the cluster. | 
| zone_id | string | X | √ | The ID of the zone where the cluster is deployed. | 
| cluster_healthy | string | X | √ | The health status of the cluster. | 
| current_version | string | X | √ | The version of the cluster. | 
| external_loadbalancer_id | string | X | √ | The ID of the Server Load Balancer (SLB) instance deployed in the cluster. | 
| node_status | string | X | √ | The status of cluster nodes. | 
| network_mode | string | X | √ | The network type of the cluster. | 
| parameters | string | X | √ |  | 
| cluster_id | string | X | √ | The ID of the cluster. | 
| size | int | X | √ | The number of nodes in the cluster. | 
| data_disk_size | int | X | √ | The size of a data disk. | 
| maintenance_info | string | X | √ |  | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the cluster. | 
| docker_version | string | X | √ | The version of Docker. | 
| next_version | string | X | √ |  | 
| updated | timestamp | X | √ | The time when the cluster was updated. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| state | string | X | √ | The status of the cluster. | 
| outputs | string | X | √ |  | 
| profile | string | X | √ | The identifier of the cluster. | 
| resource_group_id | string | X | √ | The ID of the resource group to which the cluster belongs. | 
| vpc_id | string | X | √ | The ID of the VPC used by the cluster. | 
| title | string | X | √ | Title of the resource. | 
| enabled_migration | string | X | √ |  | 
| init_version | string | X | √ | The initial version of the cluster. | 
| need_update_agent | string | X | √ |  | 
| private_zone | string | X | √ | Indicates whether PrivateZone is enabled for the cluster. | 
| maintenance_window | json | X | √ |  | 
| master_url | json | X | √ | The endpoints that are open for connections to the cluster. | 
| cluster_namespace | json | X | √ |  | 
| tags | json | X | √ | A map of tags for the resource. | 
| capabilities | string | X | √ |  | 
| cluster_type | string | X | √ | The type of the cluster. | 
| data_disk_category | string | X | √ | The type of data disks. | 
| upgrade_components | string | X | √ |  | 
| meta_data | json | X | √ | The metadata of the cluster. | 
| tags_src | json | X | √ | A list of tags attached with the cluster. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| created_at | timestamp | X | √ | The time when the cluster was created. | 
| instance_type | string | X | √ | The Elastic Compute Service (ECS) instance type of cluster nodes. | 
| subnet_cidr | cidr | X | √ | The CIDR block of pods in the cluster. | 
| vswitch_cidr | cidr | X | √ | The CIDR block of VSwitches. | 
| port | string | X | √ | Container port in Kubernetes. | 
| swarm_mode | string | X | √ |  | 
| cluster_log | json | X | √ | The logs of a cluster. | 
| name | string | X | √ | The name of the cluster. | 
| cluster_spec | string | X | √ |  | 
| deletion_protection | bool | X | √ | Indicates whether deletion protection is enabled for the cluster. | 
| gw_bridge | string | X | √ |  | 
| selefra_id | string | X | √ | primary keys value md5 | 


