resource "routeros_interface_ethernet_switch_qos_priority_flow_control" "control" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
