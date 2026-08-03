resource "routeros_interface_bridge_host" "host" {
  bridge      = "example"
  interface   = "ether1"
  mac_address = "00:00:5E:00:53:01"
}
