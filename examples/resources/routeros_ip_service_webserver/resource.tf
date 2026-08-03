resource "routeros_ip_service_webserver" "webserver" {
  acme_plain   = true
  crl_plain    = true
  graphs_plain = true
}
