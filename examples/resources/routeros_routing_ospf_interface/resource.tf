resource "routeros_routing_ospf_interface" "interface" {
  area      = "example"
  interface = "ether1"
}
