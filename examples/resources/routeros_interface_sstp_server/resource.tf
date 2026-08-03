resource "routeros_interface_sstp_server" "server" {
  certificate     = "example"
  ciphers         = "null"
  default_profile = "example"
}
