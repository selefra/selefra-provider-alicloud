package provider

import (
	"context"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
	"github.com/selefra/selefra-provider-alicloud/constants"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"github.com/spf13/viper"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Alicloud,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				// 1. first try read config from selefra config
				accessKey := config.GetString("providers.0.accounts.access_key")
				secretKey := config.GetString("providers.0.accounts.secret_key")
				regions := config.GetStringSlice("providers.0.accounts.regions")
				var alicloudConfig *alicloud_client.AliCloudConfig
				if accessKey != constants.Constants_18 || secretKey != constants.Constants_19 || len(regions) != 0 {
					alicloudConfig = &alicloud_client.AliCloudConfig{
						AccessKey: pointer.ToStringPointerOrNilIfEmpty(accessKey),
						SecretKey: pointer.ToStringPointerOrNilIfEmpty(secretKey),
						Regions:   regions,
					}
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
				return `#accounts:
  # Authenticate with the 'access_key' and 'secret_key' arguments.
#  access_key: "xxx"
#  secret_key: "xxx"
  # Optional. By default Selefra requires region as part of credentials, pass specific regions as environment variables 'ALIBABACLOUD_REGION_ID', 'ALICLOUD_REGION_ID' or 'ALICLOUD_REGION'.
#  regions:
#    - "us-east-1"
#    - "ap-south-1"`
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
