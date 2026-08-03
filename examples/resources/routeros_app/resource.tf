resource "routeros_app" "app" {
  auto_update             = true
  container_command_lines = "example"
  devices                 = "example"
}
