resource "routeros_mpls_traffic_eng_interface" "interface" {
  interface         = "ether1"
  bandwidth         = "example"
  blockade_k_factor = "example"
}
