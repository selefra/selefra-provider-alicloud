# Table: alicloud_security_center_field_statistics

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| general_asset_count | int | X | √ | The number of general assets. | 
| important_asset_count | int | X | √ | The number of important assets. | 
| offline_instance_count | int | X | √ | The number of offline servers. | 
| unprotected_instance_count | int | X | √ | The number of unprotected assets. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| new_instance_count | int | X | √ | The number of new servers. | 
| region_count | int | X | √ | The number of regions to which the servers belong. | 
| category_count | int | X | √ | The number of assets category. | 
| group_count | int | X | √ | The number of asset groups. | 
| instance_count | int | X | √ | The total number of assets of the specified type. | 
| not_running_status_count | int | X | √ | The number of inactive servers. | 
| risk_instance_count | int | X | √ | The number of assets that are at risk. | 
| test_asset_count | int | X | √ | The number of test assets. | 
| vpc_count | int | X | √ | The number of VPCs. | 
| selefra_id | string | X | √ | primary keys value md5 | 


