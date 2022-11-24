package main

import (
	"github.com/selefra/selefra-provider-sdk/grpc/serve"
	"github.com/selefra/selefra-provider-alicloud/provider"
)

func main() {

	myProvider := provider.GetProvider()
	serve.Serve(myProvider.Name, myProvider)

}
