resource "routeros_tool_romon_port" "port" {
  interface = "ether1"
  comment   = "Managed by OpenTofu"
  cost      = 1
}
