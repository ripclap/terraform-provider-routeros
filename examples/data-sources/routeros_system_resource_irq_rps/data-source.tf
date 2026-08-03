# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_resource_irq_rps" "example" {
  filter = {
    name = "example"
  }
}

output "system_resource_irq_rps" {
  value = data.routeros_system_resource_irq_rps.example.entries[*].name
}
