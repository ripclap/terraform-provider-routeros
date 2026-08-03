resource "routeros_routing_pimsm_interface_template" "template" {
  instance    = "example"
  disabled    = true
  hello_delay = "10s"
}
