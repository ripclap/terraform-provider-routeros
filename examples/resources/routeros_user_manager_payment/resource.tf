resource "routeros_user_manager_payment" "payment" {
  user     = "example"
  currency = "example"
  method   = "authorize-net"
}
