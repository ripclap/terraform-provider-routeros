resource "routeros_routing_isis_instance" "instance" {
  name      = "example"
  afi       = "ip"
  areas_max = 1
}
