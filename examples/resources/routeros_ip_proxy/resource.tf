resource "routeros_ip_proxy" "proxy" {
  always_from_cache   = true
  anonymous           = true
  cache_administrator = "example"
}
