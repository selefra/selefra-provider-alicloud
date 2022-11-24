# Table: alicloud_cs_kubernetes_cluster_node

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ | The states of the nodes in the node pool. | 
| creation_time | timestamp | X | √ | The time when the node was created. | 
| expired_time | timestamp | X | √ | The expiration time of the node. | 
| instance_name | string | X | √ | The name of the node. This name contains the ID of the cluster to which the node is deployed. | 
| instance_status | string | X | √ | The state of the node. | 
| instance_type | string | X | √ | The instance type of the node. | 
| ip_address | string | X | √ | The IP address of the node. | 
| cluster_id | string | X | √ | The ID of the cluster that the node pool belongs to. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| image_id | string | X | √ | The ID of the system image that is used by the node. | 
| is_aliyun_node | bool | X | √ | Indicates whether the instance is provided by Alibaba Cloud. | 
| source | string | X | √ | Indicates how the nodes in the node pool were initialized. The nodes can be manually created or created by using Resource Orchestration Service (ROS). | 
| instance_id | string | X | √ | The ID of the ECS instance. | 
| instance_role | string | X | √ | The role of the node. | 
| instance_charge_type | string | X | √ | The billing method of the node. | 
| nodepool_id | string | X | √ | The ID of the node pool. | 
| title | string | X | √ | Title of the resource. | 
| node_name | string | X | √ | The name of the node in the ACK cluster. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| instance_type_family | string | X | √ | The ECS instance family of the node. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| node_status | string | X | √ | Indicates whether the node is ready in the ACK cluster. Valid values: true, false. | 
| error_message | string | X | √ | The error message generated when the node was created. | 
| host_name | string | X | √ | The name of the host. | 


