resource "routeros_routing_igmp_proxy" "proxy" {
  query_interval          = "10s"
  query_response_interval = "10s"
  quick_leave             = true
}
