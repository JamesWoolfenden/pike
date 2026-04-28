resource "azurerm_netapp_volume_group_sap_hana" "pike_gen" {
  name                   = "${var.prefix}-netapp-volumegroup"
  location               = "pike"
  resource_group_name    = "pike"
  account_name           = "pike"
  group_description      = "Test volume group"
  application_identifier = "TST"

  volume {
    name                         = "${var.prefix}-netapp-volume-1"
    volume_path                  = "my-unique-file-path-1"
    service_level                = "Standard"
    capacity_pool_id             = "pike"
    subnet_id                    = "pike"
    proximity_placement_group_id = "pike"
    volume_spec_name             = "data"
    storage_quota_in_gb          = 1024
    throughput_in_mibps          = 24
    protocols                    = ["NFSv4.1"]
    security_style               = "unix"
    snapshot_directory_visible   = false

    export_policy_rule {
      rule_index          = 1
      allowed_clients     = "0.0.0.0/0"
      nfsv3_enabled       = false
      nfsv41_enabled      = true
      unix_read_only      = false
      unix_read_write     = true
      root_access_enabled = false
    }

    tags = {
      "foo" = "bar"
    }
  }

  volume {
    name                         = "${var.prefix}-netapp-volume-2"
    volume_path                  = "my-unique-file-path-2"
    service_level                = "Standard"
    capacity_pool_id             = "pike"
    subnet_id                    = "pike"
    proximity_placement_group_id = "pike"
    volume_spec_name             = "log"
    storage_quota_in_gb          = 1024
    throughput_in_mibps          = 24
    protocols                    = ["NFSv4.1"]
    security_style               = "unix"
    snapshot_directory_visible   = false

    export_policy_rule {
      rule_index          = 1
      allowed_clients     = "0.0.0.0/0"
      nfsv3_enabled       = false
      nfsv41_enabled      = true
      unix_read_only      = false
      unix_read_write     = true
      root_access_enabled = false
    }

    tags = {
      "foo" = "bar"
    }
  }

  volume {
    name                         = "${var.prefix}-netapp-volume-3"
    volume_path                  = "my-unique-file-path-3"
    service_level                = "Standard"
    capacity_pool_id             = "pike"
    subnet_id                    = "pike"
    proximity_placement_group_id = "pike"
    volume_spec_name             = "shared"
    storage_quota_in_gb          = 1024
    throughput_in_mibps          = 24
    protocols                    = ["NFSv4.1"]
    security_style               = "unix"
    snapshot_directory_visible   = false

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

  depends_on = [
    azurerm_linux_virtual_machine.example,
    azurerm_proximity_placement_group.example
  ]
}
