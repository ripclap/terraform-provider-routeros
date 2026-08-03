# Every entry in the menu.
data "routeros_certificate_crl" "example" {}

output "certificate_crl" {
  value = data.routeros_certificate_crl.example.entries[*].id
}
