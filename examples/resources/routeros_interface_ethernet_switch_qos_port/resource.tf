resource "routeros_interface_ethernet_switch_qos_port" "port" {
  egress_rate_queue0 = "example"
  egress_rate_queue1 = "example"
  egress_rate_queue2 = "example"
}
