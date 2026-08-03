resource "routeros_openflow_port" "port" {
  interface = "ether1"
  switch    = "example"
}
