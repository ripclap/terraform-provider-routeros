resource "routeros_tool_traffic_generator_port" "port" {
  interface = "ether1"
  name      = "example"
}
