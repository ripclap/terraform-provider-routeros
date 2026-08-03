resource "routeros_interface_wifi_network" "network" {
  beacon_interval = "10s"
  comment         = "Managed by OpenTofu"
  disabled        = true
}
