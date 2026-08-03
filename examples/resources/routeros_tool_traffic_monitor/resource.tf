resource "routeros_tool_traffic_monitor" "monitor" {
  interface = "ether1"
  name      = "example"
}
