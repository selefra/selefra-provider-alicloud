package alicloud_client

type AliCloudClient struct {
	Region		string
	AliCloudConfig	*AliCloudConfig
}

type AliCloudConfig struct {
	Regions		[]string	`cty:"regions"`
	AccessKey	*string		`cty:"access_key"`
	SecretKey	*string		`cty:"secret_key"`
}

func (x *AliCloudClient) CopyWithRegion(region string) *AliCloudClient {
	return &AliCloudClient{
		Region:		region,
		AliCloudConfig:	x.AliCloudConfig,
	}
}
