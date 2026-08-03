# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_iot_mqtt_brokers" "example" {
  filter = {
    name = "example"
  }
}

output "iot_mqtt_brokers" {
  value = data.routeros_iot_mqtt_brokers.example.entries[*].name
}
