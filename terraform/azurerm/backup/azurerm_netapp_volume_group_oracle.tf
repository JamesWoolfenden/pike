resource "azurerm_netapp_volume_group_oracle" "pike_gen" {
  name                   = "${var.prefix}-NetAppVolumeGroupOracle"
  location               = "pike"
  resource_group_name    = "pike"
  account_name           = "pike"
  group_description      = "Example volume group for Oracle"
  application_identifier = "TST"

  volume {
    name                       = "${var.prefix}-volume-ora1"
    volume_path                = "${var.prefix}-my-unique-file-ora-path-1"
    service_level              = "Standard"
    capacity_pool_id           = "pike"
    subnet_id                  = "pike"
    zone                       = "1"
    volume_spec_name           = "ora-data1"
    storage_quota_in_gb        = 1024
    throughput_in_mibps        = 24
    protocols                  = ["NFSv4.1"]
    security_style             = "unix"
    snapshot_directory_visible = false

    export_policy_rule {
      rule_index          = 1
      allowed_clients     = "0.0.0.0/0"
      nfsv3_enabled       = false
      nfsv41_enabled      = true
      unix_read_only      = false
      unix_read_write     = true
      root_access_enabled = false
    }
  }

  volume {
    name                       = "${var.prefix}-volume-oraLog"
    volume_path                = "${var.prefix}-my-unique-file-oralog-path"
    service_level              = "Standard"
    capacity_pool_id           = "pike"
    subnet_id                  = "pike"
    zone                       = "1"
    volume_spec_name           = "ora-log"
    storage_quota_in_gb        = 1024
    throughput_in_mibps        = 24
    protocols                  = ["NFSv4.1"]
    security_style             = "unix"
    snapshot_directory_visible = false

    export_policy_rule {
      rule_index          = 1
      allowed_clients     = "0.0.0.0/0"
      nfsv3_enabled       = false
      nfsv41_enabled      = true
      unix_read_only      = false
      unix_read_write     = true
      root_access_enabled = false
    }
  }
}
