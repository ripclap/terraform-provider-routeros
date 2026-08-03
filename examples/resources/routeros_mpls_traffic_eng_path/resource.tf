resource "routeros_mpls_traffic_eng_path" "path" {
  name                 = "example"
  affinity_exclude     = "1"
  affinity_include_all = "1"
}
