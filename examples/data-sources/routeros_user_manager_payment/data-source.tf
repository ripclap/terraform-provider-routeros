# Every entry in the menu.
data "routeros_user_manager_payment" "example" {}

output "user_manager_payment" {
  value = data.routeros_user_manager_payment.example.entries[*].id
}
