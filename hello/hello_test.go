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
          resource.TestCheckResourceAttr("data.hello.world", "message", "Hello world"),
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
