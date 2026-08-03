# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_cloud_back_to_home_file" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_cloud_back_to_home_file" {
  value = data.routeros_ip_cloud_back_to_home_file.example.entries[*].id
}
