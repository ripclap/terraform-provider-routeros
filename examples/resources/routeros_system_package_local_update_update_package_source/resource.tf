resource "routeros_system_package_local_update_update_package_source" "source" {
  address  = "192.0.2.1"
  password = "changeme"
  user     = "example"
}
