resource "routeros_routing_igmp_proxy_mfc" "mfc" {
  group    = "IPv4"
  disabled = true
  source   = "example"
}
