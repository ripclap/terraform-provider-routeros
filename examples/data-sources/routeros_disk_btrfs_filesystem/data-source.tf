# Every entry in the menu.
data "routeros_disk_btrfs_filesystem" "example" {}

output "disk_btrfs_filesystem" {
  value = data.routeros_disk_btrfs_filesystem.example.entries[*].id
}
