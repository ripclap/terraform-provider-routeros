resource "routeros_file_sync" "sync" {
  comment    = "Managed by OpenTofu"
  disabled   = true
  local_path = "example"
}
