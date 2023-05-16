# Table: alicloud_cas_certificate

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| start_date | timestamp | X | √ | The issuance date of the certificate. | 
| city | string | X | √ | The city where the organization that purchases the certificate is located. | 
| cert | string | X | √ | The certificate content, in PEM format. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| name | string | X | √ | The name of the certificate. | 
| id | float | X | √ | The ID of the certificate. | 
| buy_in_aliyun | bool | X | √ | Indicates whether the certificate was purchased from Alibaba Cloud. | 
| fingerprint | string | X | √ | The certificate fingerprint. | 
| org_name | string | X | √ | The name of the organization that purchases the certificate. | 
| end_date | timestamp | X | √ | The expiration date of the certificate. | 
| title | string | X | √ | Title of the resource. | 
| issuer | string | X | √ | The certificate authority. | 
| common | string | X | √ | The common name (CN) attribute of the certificate. | 
| province | string | X | √ | The province where the organization that purchases the certificate is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| expired | bool | X | √ | Indicates whether the certificate has expired. | 
| sans | string | X | √ | All domain names bound to the certificate. | 
| country | string | X | √ | The country where the organization that purchases the certificate is located. | 
| key | string | X | √ | The private key of the certificate, in PEM format. | 
| selefra_id | string | X | √ | primary keys value md5 | 


