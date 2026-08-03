# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_zerotier_controller_member" "example" {
  filter = {
    name = "example"
  }
}

output "zerotier_controller_member" {
  value = data.routeros_zerotier_controller_member.example.entries[*].name
}
