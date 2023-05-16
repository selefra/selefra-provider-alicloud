# Table: alicloud_ecs_instance_metric_cpu_utilization_daily

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| minimum | float | X | √ | The minimum metric value for the data point. | 
| timestamp | timestamp | X | √ | The timestamp used for the data point. | 
| instance_id | string | X | √ | The ID of the instance. | 
| metric_name | string | X | √ | The name of the metric. | 
| namespace | string | X | √ | The metric namespace. | 
| average | float | X | √ | The average of the metric values that correspond to the data point. | 
| maximum | float | X | √ | The maximum metric value for the data point. | 
| selefra_id | string | X | √ | primary keys value md5 | 
| alicloud_ecs_instance_selefra_id | string | X | √ | fk to alicloud_ecs_instance.selefra_id | 


