# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_capsman_provisioning" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "capsman_provisioning" {
  value = data.routeros_capsman_provisioning.example.entries[*].id
}
