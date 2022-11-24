# Table: alicloud_oss_bucket

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| storage_class | string | X | √ | The storage class of objects in the bucket. | 
| lifecycle_rules | json | X | √ | A list of lifecycle rules for a bucket. | 
| logging | json | X | √ | Indicates the container used to store access logging configuration of a bucket. | 
| akas | json | X | √ | Array of globally unique identifier strings (also known as) for the resource. | 
| location | string | X | √ | Location of the Bucket. | 
| redundancy_type | string | X | √ | The type of disaster recovery for a bucket. Valid values: LRS and ZRS | 
| acl | string | X | √ | The access control list setting for bucket. Valid values: public-read-write, public-read, and private. public-read-write: Any users, including anonymous users can read and write objects in the bucket. Exercise caution when you set the ACL of a bucket to public-read-write. public-read: Only the owner or authorized users of this bucket can write objects in the bucket. Other users, including anonymous users can only read objects in the bucket. Exercise caution when you set the ACL of a bucket to public-read. private: Only the owner or authorized users of this bucket can read and write objects in the bucket. Other users, including anonymous users cannot access the objects in the bucket without authorization. | 
| server_side_encryption | json | X | √ | The server-side encryption configuration for bucket | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The Alicloud region in which the resource is located. | 
| account_id | string | X | √ | The Alicloud Account ID in which the resource is located. | 
| arn | string | X | √ | The Alibaba Cloud Resource Name (ARN) of the OSS bucket. | 
| policy | json | X | √ | Allows you to grant permissions on OSS resources to RAM users from your Alibaba Cloud and other Alibaba Cloud accounts. You can also control access based on the request source. | 
| name | string | X | √ | Name of the Bucket. | 
| creation_date | timestamp | X | √ | Date when the bucket was created. | 
| versioning | string | X | √ | The status of versioning for the bucket. Valid values: Enabled and Suspended. | 
| tags_src | json | X | √ | A list of tags assigned to bucket | 
| title | string | X | √ | Title of the resource. | 


