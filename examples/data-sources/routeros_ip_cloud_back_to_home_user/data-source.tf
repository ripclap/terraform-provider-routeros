# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_cloud_back_to_home_user" "example" {
  filter = {
    name = "example"
  }
}

output "ip_cloud_back_to_home_user" {
  value = data.routeros_ip_cloud_back_to_home_user.example.entries[*].name
}
