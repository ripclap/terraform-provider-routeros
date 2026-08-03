resource "routeros_interface_ethernet_switch_l3hw_settings_advanced" "advanced" {
  neigh_discovery_burst_delay = "example"
  neigh_discovery_burst_limit = 1
  neigh_discovery_interval    = "10s"
}
