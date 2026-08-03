resource "routeros_interface_macsec_profile" "profile" {
  name            = "example"
  ciphers         = "aes-gcm-128"
  server_priority = 1
}
