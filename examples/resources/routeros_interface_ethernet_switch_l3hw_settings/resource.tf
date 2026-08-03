resource "routeros_interface_ethernet_switch_l3hw_settings" "settings" {
  autorestart         = true
  fasttrack_hw        = true
  icmp_reply_on_error = true
}
