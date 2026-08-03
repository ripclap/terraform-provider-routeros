resource "routeros_routing_rip_interface_template" "template" {
  instance = "example"
  cost     = 1
  disabled = true
}
