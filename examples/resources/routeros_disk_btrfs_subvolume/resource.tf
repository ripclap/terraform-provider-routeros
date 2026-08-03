resource "routeros_disk_btrfs_subvolume" "subvolume" {
  fs   = "example"
  name = "example"
}
