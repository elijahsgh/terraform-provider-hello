package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceHello(t *testing.T) {
  resource.UnitTest(t, resource.TestCase{
    PreCheck:  func() { testAccPreCheck(t) },
    Providers: testAccProviders,
    Steps: []resource.TestStep{
      {
        Config: testAccResourceHelloConfig,
        Check: resource.ComposeTestCheckFunc(
          testAccResourceHello_string("hello.world", "hello, world"),
          resource.TestCheckResourceAttr("hello.world", "message", "hello, world"),
        ),
      },
    },
  })
}

func testAccResourceHello_string(id string, want string) resource.TestCheckFunc {
  return func(s *terraform.State) error {
    rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found: %s", id)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

    return nil
  }
}

const (
	testAccResourceHelloConfig = `
data "hello" "world" { }
`
)
