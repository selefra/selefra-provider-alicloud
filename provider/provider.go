package provider

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	pvtz20180101 "github.com/alibabacloud-go/pvtz-20180101/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/constants"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Alicloud,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				// 1. first try read config from selefra config
				//accessKey := config.GetString("providers.0.accounts.access_key")
				//secretKey := config.GetString("providers.0.accounts.secret_key")
				//regions := config.GetStringSlice("providers.0.accounts.regions")
				accessKey := config.GetString("access_key")
				secretKey := config.GetString("secret_key")

				var alicloudConfig *alicloud_client.AliCloudConfig
				if accessKey == constants.Constants_18 || secretKey == constants.Constants_19 {
					var err error
					accessKey, secretKey, err = alicloud_client.GetEnv(ctx, clientMeta, nil, nil)
					if err != nil {
						return nil, diagnostics.AddErrorMsg("create alicloud error: %s", err.Error())
					}
				}

				// regions init
				regions := config.GetStringSlice("accounts.regions")
				if len(regions) == 0 {
					regionsString, exists := os.LookupEnv("ALIBABACLOUD_REGIONS")
					if exists && regionsString != "" {
						for _, region := range strings.Split(regionsString, ",") {
							regions = append(regions, strings.TrimSpace(region))
						}
					}
					//} else {
					//	region, exists := os.LookupEnv("ALIBABACLOUD_REGION_ID")
					//	if exists && region != "" {
					//		regions = append(regions, region)
					//	} else {
					//		region, exists := os.LookupEnv("ALICLOUD_REGION_ID")
					//		if exists && region != "" {
					//			regions = append(regions, region)
					//		} else {
					//			region, exists := os.LookupEnv("REGION")
					//			if exists && region != "" {
					//				regions = append(regions, region)
					//			}
					//		}
					//	}
					//}
				}

				// ------------------------------------------------- --------------------------------------------------------------------

				client, err := CreateClient(tea.String(accessKey), tea.String(secretKey))
				if err != nil {
					return nil, diagnostics.AddErrorMsg("create alicloud error: %s", err.Error())
				}

				describeRegionsRequest := &pvtz20180101.DescribeRegionsRequest{}
				response, err := func() (response *pvtz20180101.DescribeRegionsResponse, err error) {
					defer func() {
						if r := tea.Recover(recover()); r != nil {
							err = r
						}
					}()
					for tryTimes := 0; tryTimes < 3; tryTimes++ {
						response, err = client.DescribeRegions(describeRegionsRequest)
						if err != nil {
							continue
						}
						return response, nil
					}
					return
				}()

				if err != nil {
					return nil, diagnostics.AddErrorMsg("init client error: %s", err.Error())
				} else {
					latestRegionSet := make(map[string]struct{}, 0)
					if response.Body != nil && response.Body.Regions != nil && len(response.Body.Regions.Region) != 0 {
						for _, region := range response.Body.Regions.Region {
							if region.RegionId != nil {
								regionId := *region.RegionId
								//if strings.Contains(regionId, "-gov-") || strings.Contains(regionId, "-finance-") {
								//	continue
								//}
								latestRegionSet[regionId] = struct{}{}
							}
						}
					}
					if len(latestRegionSet) == 0 {
						clientMeta.DebugF("get latest regions error, client init error")
						return nil, diagnostics.AddErrorMsg("get latest regions error, client init error")
					}
					// merge latest region
					for regionId, _ := range latestRegionSet {
						alicloud_client.AlicloudRegions[regionId] = struct{}{}
					}
					clientMeta.DebugF("init all latest regions: ", zap.Any("latestRegions", latestRegionSet))
				}

				// ------------------------------------------------- --------------------------------------------------------------------

				// Set the variable so that it can be retrieved later
				_ = os.Setenv("ALIBABACLOUD_ACCESS_KEY_ID", accessKey)
				_ = os.Setenv("ALIBABACLOUD_ACCESS_KEY_SECRET", secretKey)

				alicloudConfig = &alicloud_client.AliCloudConfig{
					AccessKey: pointer.ToStringPointerOrNilIfEmpty(accessKey),
					SecretKey: pointer.ToStringPointerOrNilIfEmpty(secretKey),
					Regions:   regions,
				}

				return []any{
					&alicloud_client.AliCloudClient{
						Region:         constants.Cnbeijing,
						AliCloudConfig: alicloudConfig,
					},
				}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# Authenticate with the 'access_key' and 'secret_key' arguments.
access_key: "xxx"
secret_key: "xxx"
# Optional. By default Selefra requires region as part of credentials, pass specific regions as environment variables 'ALIBABACLOUD_REGION_ID', 'ALICLOUD_REGION_ID' or 'ALICLOUD_REGION'.
regions:
  - "us-east-1"
  - "ap-south-1"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.LOGINDISABLED,
				constants.NA,
				constants.Constants_20,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *pvtz20180101.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("pvtz.aliyuncs.com")
	_result = &pvtz20180101.Client{}
	_result, _err = pvtz20180101.NewClient(config)
	return _result, _err
}
