resource "routeros_routing_isis_interface_template" "template" {
  instance                = "example"
  bcast_l1_csnp_interval  = "10s"
  bcast_l1_hello_interval = "10s"
}
