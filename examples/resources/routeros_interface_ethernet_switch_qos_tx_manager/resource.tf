resource "routeros_interface_ethernet_switch_qos_tx_manager" "manager" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
