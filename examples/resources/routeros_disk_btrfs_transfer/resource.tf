resource "routeros_disk_btrfs_transfer" "transfer" {
  type = "receive"
  file = "url"
  fs   = "example"
}
