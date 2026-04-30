package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nokia/eda/apps/terraform-provider-connect/internal/provider"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "github.com/nokia-eda/connect-v2",
	}

	err := providerserver.Serve(context.Background(), provider.New("v2"), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
