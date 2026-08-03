resource "routeros_interface_bridge_mdb" "mdb" {
  bridge = "example"
  group  = "IPv4"
}
