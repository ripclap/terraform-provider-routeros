resource "routeros_ppp_l2tp_secret" "secret" {
  address = "192.0.2.1"
  comment = "Managed by OpenTofu"
  secret  = "changeme"
}
