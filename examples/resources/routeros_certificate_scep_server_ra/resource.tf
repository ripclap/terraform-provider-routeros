resource "routeros_certificate_scep_server_ra" "ra" {
  name               = "example"
  ca_identity        = "example"
  challenge_password = "changeme"
}
