resource "routeros_disk_btrfs_filesystem" "filesystem" {
  uuid              = "example"
  default_subvolume = "example"
  label             = "example"
}
