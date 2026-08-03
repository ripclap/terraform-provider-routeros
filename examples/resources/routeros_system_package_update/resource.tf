resource "routeros_system_package_update" "update" {
  channel           = "development"
  check_certificate = "no"
  ip_version        = "auto"
}
