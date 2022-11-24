package provider

import (
	"github.com/selefra/selefra-provider-alicloud/constants"
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/spf13/viper"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:		constants.Alicloud,
		Version:	Version,
		TableList:	GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				accessKey := config.GetString(constants.Accesskey)
				secretKey := config.GetString(constants.Secretkey)
				regisons := config.GetStringSlice(constants.Regions)

				var alicloudConfig *alicloud_client.AliCloudConfig
				if accessKey != constants.Constants_18 || secretKey != constants.Constants_19 || len(regisons) != 0 {
					alicloudConfig = &alicloud_client.AliCloudConfig{
						AccessKey:	pointer.ToStringPointerOrNilIfEmpty(accessKey),
						SecretKey:	pointer.ToStringPointerOrNilIfEmpty(secretKey),
						Regions:	regisons,
					}
				}
				return []any{
					&alicloud_client.AliCloudClient{
						Region:		constants.Cnbeijing,
						AliCloudConfig:	alicloudConfig,
					},
				}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `accounts:
  # Authenticate with the 'access_key' and 'secret_key' arguments.
  # access_key: "xxx"
  # secret_key: "xxx"
  # Optional. By default Selefra requires region as part of credentials, pass specific regions as environment variables 'ALIBABACLOUD_REGION_ID', 'ALICLOUD_REGION_ID' or 'ALICLOUD_REGION'. 
  # regions:
  #  - "us-east-1"
  #  - "ap-south-1"`
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
			DataSourcePullResultAutoExpand:	true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
