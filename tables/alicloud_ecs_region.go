package tables

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-alicloud/table_schema_generator"
	"github.com/selefra/selefra-provider-alicloud/alicloud_client"
)

type TableAlicloudEcsRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlicloudEcsRegionGenerator{}

func (x *TableAlicloudEcsRegionGenerator) GetTableName() string {
	return "alicloud_ecs_region"
}

func (x *TableAlicloudEcsRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlicloudEcsRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlicloudEcsRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlicloudEcsRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := alicloud_client.ECSRegionService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			request := ecs.CreateDescribeRegionsRequest()
			request.Scheme = "https"
			request.AcceptLanguage = "en-US"

			response, err := client.DescribeRegions(request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range response.Regions.Region {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getRegionAkas(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	data := result.(ecs.Region)

	commonData, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ecs::" + accountID + ":region/" + data.RegionId}, nil
}

func (x *TableAlicloudEcsRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlicloudEcsRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Alicloud region in which the resource is located.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("local_name").ColumnType(schema.ColumnTypeString).Description("The local name of the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_endpoint").ColumnType(schema.ColumnTypeString).Description("The endpoint of the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Indicates whether the cluster is sold out.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).Description("Array of globally unique identifier strings (also known as) for the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := getRegionAkas(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("RegionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The Alicloud Account ID in which the resource is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("AccountID")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableAlicloudEcsRegionGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlicloudEcsZoneGenerator{}),
	}
}
