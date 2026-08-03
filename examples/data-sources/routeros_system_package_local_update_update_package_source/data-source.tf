# Every entry in the menu.
data "routeros_system_package_local_update_update_package_source" "example" {}

output "system_package_local_update_update_package_source" {
  value = data.routeros_system_package_local_update_update_package_source.example.entries[*].id
}
