package hello

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
//		ResourcesMap: map[string]*schema.Resource{
		DataSourcesMap: map[string]*schema.Resource{
			"hello": dataSourceHello(),
		},
	}
}
