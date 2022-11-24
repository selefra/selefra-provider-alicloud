# Table: alicloud_cas_certificate

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| issuer | string | X | √ | The certificate authority. | 
| buy_in_aliyun | bool | X | √ | Indicates whether the certificate was purchased from Alibaba Cloud. | 
| fingerprint | string | X | √ | The certificate fingerprint. | 
| title | string | X | √ | Title of the resource. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| end_date | timestamp | X | √ | The expiration date of the certificate. | 
| sans | string | X | √ | All domain names bound to the certificate. | 
| province | string | X | √ | The province where the organization that purchases the certificate is located. | 
| city | string | X | √ | The city where the organization that purchases the certificate is located. | 
| cert | string | X | √ | The certificate content, in PEM format. | 
| key | string | X | √ | The private key of the certificate, in PEM format. | 
| name | string | X | √ | The name of the certificate. | 
| org_name | string | X | √ | The name of the organization that purchases the certificate. | 
| common | string | X | √ | The common name (CN) attribute of the certificate. | 
| expired | bool | X | √ | Indicates whether the certificate has expired. | 
| id | float | X | √ | The ID of the certificate. | 
| start_date | timestamp | X | √ | The issuance date of the certificate. | 
| country | string | X | √ | The country where the organization that purchases the certificate is located. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 


