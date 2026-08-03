resource "routeros_interface_wifi_network_radio" "radio" {
  comment      = "Managed by OpenTofu"
  disabled     = true
  extra_labels = "example"
}
