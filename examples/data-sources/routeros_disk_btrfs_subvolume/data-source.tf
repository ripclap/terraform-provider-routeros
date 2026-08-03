# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_disk_btrfs_subvolume" "example" {
  filter = {
    name = "example"
  }
}

output "disk_btrfs_subvolume" {
  value = data.routeros_disk_btrfs_subvolume.example.entries[*].name
}
