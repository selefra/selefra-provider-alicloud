# Table: alicloud_rds_instance_metric_connections_daily

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ | The metric namespace. | 
| average | float | X | √ | The average of the metric values that correspond to the data point. | 
| maximum | float | X | √ | The maximum metric value for the data point. | 
| minimum | float | X | √ | The minimum metric value for the data point. | 
| timestamp | timestamp | X | √ | The timestamp used for the data point. | 
| db_instance_id | string | X | √ | The ID of the single instance to query. | 
| metric_name | string | X | √ | The name of the metric. | 


