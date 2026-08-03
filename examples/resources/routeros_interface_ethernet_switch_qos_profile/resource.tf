resource "routeros_interface_ethernet_switch_qos_profile" "profile" {
  name    = "example"
  automap = true
  comment = "Managed by OpenTofu"
}
