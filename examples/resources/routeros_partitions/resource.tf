resource "routeros_partitions" "partitions" {
  name        = "example"
  comment     = "Managed by OpenTofu"
  fallback_to = "example"
}
