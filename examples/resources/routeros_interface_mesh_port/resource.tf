resource "routeros_interface_mesh_port" "port" {
  interface = "ether1"
  mesh      = "example"
}
