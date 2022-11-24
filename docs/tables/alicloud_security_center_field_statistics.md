# Table: alicloud_security_center_field_statistics

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| important_asset_count | int | X | √ | The number of important assets. | 
| region_count | int | X | √ | The number of regions to which the servers belong. | 
| test_asset_count | int | X | √ | The number of test assets. | 
| group_count | int | X | √ | The number of asset groups. | 
| risk_instance_count | int | X | √ | The number of assets that are at risk. | 
| unprotected_instance_count | int | X | √ | The number of unprotected assets. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| general_asset_count | int | X | √ | The number of general assets. | 
| new_instance_count | int | X | √ | The number of new servers. | 
| not_running_status_count | int | X | √ | The number of inactive servers. | 
| offline_instance_count | int | X | √ | The number of offline servers. | 
| category_count | int | X | √ | The number of assets category. | 
| instance_count | int | X | √ | The total number of assets of the specified type. | 
| vpc_count | int | X | √ | The number of VPCs. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 


