resource "routeros_mpls_traffic_eng_tunnel" "tunnel" {
  name                 = "example"
  affinity_exclude     = "1"
  affinity_include_all = "1"
}
