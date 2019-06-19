package main

import (
        "github.com/hashicorp/terraform/plugin"
        "hello"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: hello.Provider})
}
