resource "routeros_queue_interface" "interface" {
  interface = "ether1"
  queue     = "example"
}
