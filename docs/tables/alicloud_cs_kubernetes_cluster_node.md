# Table: alicloud_cs_kubernetes_cluster_node

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_type | string | X | √ | The instance type of the node. | 
| node_name | string | X | √ | The name of the node in the ACK cluster. | 
| creation_time | timestamp | X | √ | The time when the node was created. | 
| expired_time | timestamp | X | √ | The expiration time of the node. | 
| instance_name | string | X | √ | The name of the node. This name contains the ID of the cluster to which the node is deployed. | 
| instance_role | string | X | √ | The role of the node. | 
| instance_status | string | X | √ | The state of the node. | 
| source | string | X | √ | Indicates how the nodes in the node pool were initialized. The nodes can be manually created or created by using Resource Orchestration Service (ROS). | 
| ip_address | string | X | √ | The IP address of the node. | 
| host_name | string | X | √ | The name of the host. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| nodepool_id | string | X | √ | The ID of the node pool. | 
| error_message | string | X | √ | The error message generated when the node was created. | 
| cluster_id | string | X | √ | The ID of the cluster that the node pool belongs to. | 
| is_aliyun_node | bool | X | √ | Indicates whether the instance is provided by Alibaba Cloud. | 
| instance_id | string | X | √ | The ID of the ECS instance. | 
| instance_charge_type | string | X | √ | The billing method of the node. | 
| instance_type_family | string | X | √ | The ECS instance family of the node. | 
| image_id | string | X | √ | The ID of the system image that is used by the node. | 
| state | string | X | √ | The states of the nodes in the node pool. | 
| node_status | string | X | √ | Indicates whether the node is ready in the ACK cluster. Valid values: true, false. | 
| title | string | X | √ | Title of the resource. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| selefra_id | string | X | √ | primary keys value md5 | 
| alicloud_cs_kubernetes_cluster_selefra_id | string | X | √ | fk to alicloud_cs_kubernetes_cluster.selefra_id | 


