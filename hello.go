package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func dataSourceHello() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHelloRead,
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceHelloRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("message", "Hello world")
	return nil
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return &schema.Provider{
				DataSourcesMap: map[string]*schema.Resource{
					"hello": dataSourceHello(),
				},
			}
		},
	})
}
