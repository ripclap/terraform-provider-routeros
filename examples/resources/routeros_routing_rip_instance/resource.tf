resource "routeros_routing_rip_instance" "instance" {
  name     = "example"
  afi      = "ip"
  disabled = true
}
