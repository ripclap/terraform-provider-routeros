resource "routeros_ip_ipsec_key_qkd" "qkd" {
  address     = "192.0.2.1"
  cache_size  = 1
  certificate = "example"
}
