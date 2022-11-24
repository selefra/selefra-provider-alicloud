package alicloud_client

import (
	"github.com/selefra/selefra-provider-alicloud/constants"
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"os"
	"strings"
)

func BuildRegionList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {

		client := taskClient.(*AliCloudClient)

		if client.AliCloudConfig != nil && len(client.AliCloudConfig.Regions) != 0 {
			regions := client.AliCloudConfig.Regions

			if len(getInvalidRegions(regions)) > 0 {
				panic(constants.NnConnectionconfighaveinvalidregions + strings.Join(getInvalidRegions(regions), constants.Constants_0) + constants.Edityourconnectionconfigurationfileandthenrestartselefra)
			}

			slice := make([]*schema.ClientTaskContext, 0)
			for _, region := range regions {
				slice = append(slice, &schema.ClientTaskContext{
					Task:	task.Clone(),
					Client:	client.CopyWithRegion(region),
				})
			}
			return slice
		}

		return []*schema.ClientTaskContext{
			&schema.ClientTaskContext{
				Task:	task.Clone(),
				Client:	client.CopyWithRegion(GetDefaultRegion(ctx, clientMeta, taskClient, task)),
			},
		}
	}
}

var alicloudRegions = map[string]struct{}{
	constants.Cnbeijing:		{},
	constants.Cnbeijingfinance:	{},
	constants.Cnchengdu:		{},
	constants.Cnguangzhou:		{},
	constants.Cnhangzhou:		{},
	constants.Cnheyuan:		{},
	constants.Cnhongkong:		{},
	constants.Cnhuhehaote:		{},
	constants.Cnqingdao:		{},
	constants.Cnshanghai:		{},
	constants.Cnshanghaifinance:	{},
	constants.Cnshenzhen:		{},
	constants.Cnshenzhenfinance:	{},
	constants.Cnwulanchabu:		{},
	constants.Cnzhangjiakou:		{},
	constants.Apnortheast:		{},
	constants.Apsouth:		{},
	constants.Apsoutheast:		{},
	"ap-southeast-2":	{},
	"ap-southeast-3":	{},
	"ap-southeast-5":	{},
	constants.Eucentral:	{},
	constants.Euwest:		{},
	constants.Meeast:		{},
	constants.Useast:		{},
	constants.Uswest:		{},
}

func getInvalidRegions(regions []string) []string {
	var invalidRegions []string
	for _, region := range regions {
		if _, exists := alicloudRegions[region]; exists {
			invalidRegions = append(invalidRegions, region)
		}
	}
	return invalidRegions
}

func GetDefaultRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) string {

	alicloudConfig := taskClient.(*AliCloudClient).AliCloudConfig

	var regions []string
	var region string

	if alicloudConfig != nil && alicloudConfig.Regions != nil {
		regions = alicloudConfig.Regions
	}

	if len(regions) > 0 {

		region = regions[0]

		if len(getInvalidRegions([]string{region})) > 0 {
			panic(constants.NnConnectionconfighaveinvalidregion + region + ". Edit your connection configuration file and then restart selefra")
		}
		return region
	}

	if region == constants.Constants_1 {
		region = os.Getenv(constants.ALIBABACLOUDREGIONID)
		if region == constants.Constants_2 {
			region = os.Getenv(constants.ALICLOUDREGIONID)
			if region == constants.Constants_3 {
				region = os.Getenv(constants.ALICLOUDREGION)
			}
		}
	}

	if region == constants.Constants_4 {
		region = constants.Cnbeijing
	}

	return region
}

func getEnv(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (secretKey string, accessKey string, err error) {

	alicloudConfig := taskClient.(*AliCloudClient).AliCloudConfig

	if alicloudConfig != nil && alicloudConfig.AccessKey != nil {
		accessKey = *alicloudConfig.AccessKey
	} else {
		var ok bool
		if accessKey, ok = os.LookupEnv(constants.ALIBABACLOUDACCESSKEYID); !ok {
			if accessKey, ok = os.LookupEnv(constants.ALICLOUDACCESSKEYID); !ok {
				if accessKey, ok = os.LookupEnv(constants.ALICLOUDACCESSKEY); !ok {
					panic(constants.NaccesskeymustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileandthenrestartselefra)
				}
			}
		}
	}

	if alicloudConfig != nil && alicloudConfig.SecretKey != nil {
		secretKey = *alicloudConfig.SecretKey
	} else {
		var ok bool
		if secretKey, ok = os.LookupEnv(constants.ALIBABACLOUDACCESSKEYSECRET); !ok {
			if secretKey, ok = os.LookupEnv(constants.ALICLOUDACCESSKEYSECRET); !ok {
				if secretKey, ok = os.LookupEnv(constants.ALICLOUDSECRETKEY); !ok {
					panic(constants.NsecretkeymustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileandthenrestartselefra)
				}
			}
		}
	}

	return accessKey, secretKey, nil
}
