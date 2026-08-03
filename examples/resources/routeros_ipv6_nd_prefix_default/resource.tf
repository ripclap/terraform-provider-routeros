resource "routeros_ipv6_nd_prefix_default" "default" {
  autonomous         = true
  dhcp6_pd_preferred = true
  preferred_lifetime = "10s"
}
