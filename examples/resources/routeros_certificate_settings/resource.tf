resource "routeros_certificate_settings" "settings" {
  builtin_trust_store = "all"
  crl_download        = true
  crl_store           = "ram"
}
