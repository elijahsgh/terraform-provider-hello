package hello

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceHello(t *testing.T) {
  resource.Test(t, resource.TestCase{
    PreCheck:  func() { testAccPreCheck(t) },
    Providers: testAccProviders,
    Steps: []resource.TestStep{
      {
        Config: testAccResourceHelloConfig,
        Check: resource.ComposeTestCheckFunc(
          resource.TestCheckResourceAttr("hello.world", "message", "hello, world"),
        ),
      },
    },
  })
}

const (
	testAccResourceHelloConfig = `
data "hello" "world" { }
`
)
