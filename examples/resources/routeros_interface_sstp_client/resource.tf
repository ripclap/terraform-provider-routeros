resource "routeros_interface_sstp_client" "client" {
  connect_to = "example"
  name       = "example"
  password   = "changeme"
  user       = "example"
}
