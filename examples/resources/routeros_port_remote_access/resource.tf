resource "routeros_port_remote_access" "access" {
  port    = "8728"
  channel = 1
  comment = "Managed by OpenTofu"
}
