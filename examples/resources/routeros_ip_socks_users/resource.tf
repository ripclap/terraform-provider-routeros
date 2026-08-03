resource "routeros_ip_socks_users" "users" {
  name     = "example"
  disabled = true
  only_one = true
}
