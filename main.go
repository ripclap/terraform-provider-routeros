package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ripclap/terraform-provider-routeros/routeros"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// Set by the release build through -ldflags.
var (
	version = "dev"
	commit  = "none"
)

func main() {
	var debug, showVersion bool

	// https://developer.hashicorp.com/terraform/plugin/debugging
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.BoolVar(&showVersion, "version", false, "print the provider version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("%s %s\n", version, commit)
		os.Exit(0)
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderAddr: "registry.opentofu.org/ripclap/routeros",
		ProviderFunc: routeros.NewProvider,
		Debug:        debug,
	})
}
