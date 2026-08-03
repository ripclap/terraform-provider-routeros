resource "routeros_system_package_local_update_mirror" "mirror" {
  check_interval = "10s"
  enabled        = true
  password       = "changeme"
}
