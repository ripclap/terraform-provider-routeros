resource "routeros_ip_ipsec_key_psk" "psk" {
  key    = "changeme"
  peer   = "192.0.2.1"
  psk_id = "example"
}
