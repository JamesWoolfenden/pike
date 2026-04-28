resource "azurerm_netapp_volume" "pike_gen" {
  name                       = "example-netappvolume"
  location                   = "pike"
  zone                       = "1"
  resource_group_name        = "pike"
  account_name               = "pike"
  pool_name                  = "pike"
  volume_path                = "my-unique-file-path"
  service_level              = "Premium"
  subnet_id                  = "pike"
  protocols                  = ["NFSv4.1"]
  security_style             = "unix"
  storage_quota_in_gb        = 100
  snapshot_directory_visible = false
  large_volume_enabled       = false

  # When creating volume from a snapshot
  create_from_snapshot_resource_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1"

  # Following section is only required if deploying a data protection volume (secondary)
  # to enable Cross-Region Replication feature
  data_protection_replication {
    endpoint_type             = "dst"
    remote_volume_location    = "pike"
    remote_volume_resource_id = "pike"
    replication_frequency     = "10minutes"
  }

  # Enabling Snapshot Policy for the volume
  # Note: this cannot be used in conjunction with data_protection_replication when endpoint_type is dst
  data_protection_snapshot_policy {
    snapshot_policy_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/snapshotPolicies/snapshotpolicy1"
  }

  # Enabling backup policy
  data_protection_backup_policy {
    backup_vault_id  = "pike"
    backup_policy_id = "pike"
    policy_enabled   = true
  }

  # Enabling Advanced Ransomware Protection (ARP)
  data_protection_advanced_ransomware {
    protection_enabled = true
  }

  # prevent the possibility of accidental data loss
  lifecycle {
    prevent_destroy = true
  }
}
