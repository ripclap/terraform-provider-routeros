resource "routeros_ip_dhcp_server_alert" "alert" {
  interface     = "ether1"
  alert_timeout = "10s"
  comment       = "Managed by OpenTofu"
}
