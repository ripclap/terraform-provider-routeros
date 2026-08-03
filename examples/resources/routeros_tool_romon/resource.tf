resource "routeros_tool_romon" "romon" {
  enabled  = true
  romon_id = "00:00:5E:00:53:01"
  secrets  = "changeme"
}
