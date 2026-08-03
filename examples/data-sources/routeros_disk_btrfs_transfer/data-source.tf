# Every entry in the menu.
data "routeros_disk_btrfs_transfer" "example" {}

output "disk_btrfs_transfer" {
  value = data.routeros_disk_btrfs_transfer.example.entries[*].id
}
