package alicloud_client

import (
	"context"
	"fmt"
	ims "github.com/alibabacloud-go/ims-20190815/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/actiontrail"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/selefra/selefra-provider-alicloud/constants"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func AutoscalingService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*ess.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	if region == constants.Constants_5 {
		return nil, fmt.Errorf(constants.RegionmustbepassedAutoscalingService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := ess.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func CasService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*cas.Client, error) {
	region := taskClient.(*AliCloudClient).Region
	if region == constants.Constants_6 {
		return nil, fmt.Errorf(constants.RegionmustbepassedCasService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := cas.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func CmsService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*cms.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	if region == constants.Constants_7 {
		return nil, fmt.Errorf(constants.RegionmustbepassedCmsService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := cms.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func ECSService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*ecs.Client, error) {
	region := taskClient.(*AliCloudClient).Region
	if region == constants.Constants_8 {
		return nil, fmt.Errorf(constants.RegionmustbepassedECSService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := ecs.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func ECSRegionService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*ecs.Client, error) {
	region := taskClient.(*AliCloudClient).Region
	if region == constants.Constants_9 {
		return nil, fmt.Errorf(constants.RegionmustbepassedECSRegionService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := ecs.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func IMSService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*ims.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	config := &rpc.Config{}
	config.AccessKeyId = &ak
	config.AccessKeySecret = &secret
	config.RegionId = &region

	svc, err := ims.NewClient(config)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func KMSService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*kms.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	if region == constants.Constants_10 {
		return nil, fmt.Errorf(constants.RegionmustbepassedKMSService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := kms.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func RAMService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*ram.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := ram.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func StsService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*sts.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := sts.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func VpcService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*vpc.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	if region == constants.Constants_11 {
		return nil, fmt.Errorf(constants.RegionmustbepassedVpcService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := vpc.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func OssService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*oss.Client, error) {
	region := taskClient.(*AliCloudClient).Region
	if region == constants.Constants_12 {
		return nil, fmt.Errorf(constants.RegionmustbepassedOssService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := oss.New(constants.Oss+region+constants.Aliyuncscom, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func RDSService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*rds.Client, error) {
	region := taskClient.(*AliCloudClient).Region
	if region == constants.Constants_13 {
		return nil, fmt.Errorf(constants.RegionmustbepassedRDSService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := rds.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func ActionTrailService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*actiontrail.Client, error) {
	region := taskClient.(*AliCloudClient).Region

	if region == constants.Constants_14 {
		return nil, fmt.Errorf(constants.RegionmustbepassedActionTrailService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := actiontrail.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func ContainerService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*cs.Client, error) {
	region := GetDefaultRegions(ctx, clientMeta, taskClient, task)

	if region == constants.Constants_15 {
		return nil, fmt.Errorf(constants.RegionmustbepassedContainerService)
	}
	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := cs.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func SecurityCenterService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*sas.Client, error) {
	aliCloudClient := taskClient.(*AliCloudClient)
	region := aliCloudClient.Region
	if region == constants.Constants_16 {
		return nil, fmt.Errorf(constants.RegionmustbepassedSecurityCenterService)
	}

	ak, secret, err := getEnv(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	svc, err := sas.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	return svc, nil
}
