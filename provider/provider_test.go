package provider

import (
	"github.com/selefra/selefra-provider-alicloud/constants"

	"context"
	"fmt"

	"testing"

	"github.com/selefra/selefra-provider-sdk/env"

	"github.com/selefra/selefra-provider-sdk/grpc/shard"

	"github.com/selefra/selefra-provider-sdk/provider"

	"github.com/selefra/selefra-provider-sdk/provider/schema"

	"github.com/selefra/selefra-provider-sdk/storage/database_storage/postgresql_storage"

	"github.com/selefra/selefra-utils/pkg/json_util"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

func TestProvider_PullTable(t *testing.T) {

	wk := constants.Constants_21
	config := `accounts: "test"




  # Authenticate with the 'access_key' and 'secret_key' arguments.




  # access_key: "xxx"
  # secret_key: "xxx"
  # Optional. By default Selefra requires region as part of credentials, pass specific regions as environment variables 'ALIBABACLOUD_REGION_ID', 'ALICLOUD_REGION_ID' or 'ALICLOUD_REGION'. 




  # regions:
  #  - "us-east-1"
  #  - "ap-south-1"`

	myProvider := GetProvider()

	Pull(myProvider, config, wk, constants.Constants_22)
}

func Test(t *testing.T) {

	myProvider := GetProvider()

	for _, table := range myProvider.TableList {
		fmt.Println(table.TableName)

	}

}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	initProviderRequest := &shard.ProviderInitRequest{

		Storage: &shard.Storage{

			Type: 0,

			StorageOptions: json_util.ToJsonBytes(postgresql_storage.NewPostgresqlStorageOptions(env.GetDatabaseDsn())),
		},
		Workspace: &workspace,

		IsInstallInit: pointer.TruePointer(),

		ProviderConfig: &config,
	}

	response, err := myProvider.Init(context.Background(), initProviderRequest)

	if err != nil {

		panic(diagnostics.AddFatal(constants.Initprovidererrors, err.Error()).ToString())
	}

	if diagnostics.AddDiagnostics(response.Diagnostics).HasError() {

		panic(diagnostics.ToString())

	}

	err = myProvider.PullTables(context.Background(), &shard.PullTablesRequest{

		Tables: pullTables,

		MaxGoroutines: 1,

		Timeout: 1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())

	if err != nil {

		panic(diagnostics.AddFatal(constants.Providerpulltableerrors, err.Error()).ToString())

	}

}
