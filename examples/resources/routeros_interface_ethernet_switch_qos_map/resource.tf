resource "routeros_interface_ethernet_switch_qos_map" "map" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
