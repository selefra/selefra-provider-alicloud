package main

import (
	"github.com/selefra/selefra-provider-alicloud/provider"
	"github.com/selefra/selefra-provider-sdk/grpc/serve"
)

func main() {

	myProvider := provider.GetProvider()
	serve.Serve(myProvider.Name, myProvider)

}
