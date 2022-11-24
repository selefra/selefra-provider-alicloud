package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsAutoscalingGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudSecurityCenterFieldStatisticsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsRegionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudKmsSecretGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudCasCertificateGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcNetworkAclGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcRouteTableGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsDiskGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudSecurityCenterVersionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudCmsMonitorHostGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsSnapshotGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudCsKubernetesClusterGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcEipGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudAccountGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsNetworkInterfaceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudOssBucketGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsKeyPairGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsLaunchTemplateGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcSslVpnClientCertGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsAutoProvisioningGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcSslVpnServerGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcVpnConnectionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcVpnCustomerGatewayGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRdsInstanceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamPasswordPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcNatGatewayGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudKmsKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamSecurityPreferenceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamRoleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcVpnGatewayGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsSecurityGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcVswitchGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcDhcpOptionsSetGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudVpcGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsInstanceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudRamCredentialReportGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudActionTrailGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlicloudEcsImageGenerator{}),
	}
}
