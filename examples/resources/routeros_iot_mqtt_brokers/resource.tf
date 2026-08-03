resource "routeros_iot_mqtt_brokers" "brokers" {
  name         = "example"
  address      = "192.0.2.1"
  auto_connect = true
}
