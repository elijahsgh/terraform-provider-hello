package hello

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHello() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHelloRead,
		Schema: map[string]*schema.Schema{
			"message": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The result of saying Hello to the world.",
				Elem: &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceHelloRead(d *schema.ResourceData, meta interface{}) error {
	d.SetId("hello-world")
	d.Set("message", "Hello world")
	return nil
}
