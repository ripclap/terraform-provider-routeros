resource "routeros_ip_packing" "packing" {
  interface       = "ether1"
  aggregated_size = 20
  disabled        = true
}
