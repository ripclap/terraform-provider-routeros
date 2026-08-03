resource "routeros_tool_sms" "sms" {
  allowed_number = "example"
  channel        = 1
  polling        = true
}
