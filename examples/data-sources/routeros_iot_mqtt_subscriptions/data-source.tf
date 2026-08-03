# Every entry in the menu.
data "routeros_iot_mqtt_subscriptions" "example" {}

output "iot_mqtt_subscriptions" {
  value = data.routeros_iot_mqtt_subscriptions.example.entries[*].id
}
