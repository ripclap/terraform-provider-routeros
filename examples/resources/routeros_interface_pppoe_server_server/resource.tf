resource "routeros_interface_pppoe_server_server" "server" {
  interface            = "ether1"
  accept_empty_service = true
  accept_untagged      = true
}
