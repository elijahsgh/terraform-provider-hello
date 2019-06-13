data "hello" "world" {}

output "result" {
  value = data.hello.world.message
}
