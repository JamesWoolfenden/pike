resource "azurerm_stack_hci_deployment_setting" "pike_gen" {
  stack_hci_cluster_id = "pike"
  arc_resource_ids     = [for server in "pike" : "pike"]
  version              = "10.0.0.0"

  scale_unit {
    active_directory_organizational_unit_path = "OU=hci,DC=jumpstart,DC=local"
    domain_fqdn                               = "jumpstart.local"
    secrets_location                          = azurerm_key_vault.DeploymentKeyVault.vault_uri
    name_prefix                               = "hci"

    cluster {
      azure_service_endpoint = "core.windows.net"
      cloud_account_name     = "pike"
      name                   = "pike"
      witness_type           = "Cloud"
      witness_path           = "Cloud"
    }

    host_network {
      storage_auto_ip_enabled                 = true
      storage_connectivity_switchless_enabled = false
      intent {
        name                                          = "ManagementCompute"
        adapter_property_override_enabled             = false
        qos_policy_override_enabled                   = false
        virtual_switch_configuration_override_enabled = false
        adapter = [
          "FABRIC",
          "FABRIC2",
        ]
        traffic_type = [
          "Management",
          "Compute",
        ]
        qos_policy_override {
          priority_value8021_action_cluster = "7"
          priority_value8021_action_smb     = "3"
          bandwidth_percentage_smb          = "50"
        }
        adapter_property_override {
          jumbo_packet              = "9014"
          network_direct            = "Disabled"
          network_direct_technology = "RoCEv2"
        }
      }

      intent {
        name                                          = "Storage"
        adapter_property_override_enabled             = false
        qos_policy_override_enabled                   = false
        virtual_switch_configuration_override_enabled = false
        adapter = [
          "StorageA",
          "StorageB",
        ]
        traffic_type = [
          "Storage",
        ]
        qos_policy_override {
          priority_value8021_action_cluster = "7"
          priority_value8021_action_smb     = "3"
          bandwidth_percentage_smb          = "50"
        }
        adapter_property_override {
          jumbo_packet              = "9014"
          network_direct            = "Enabled"
          network_direct_technology = "RoCEv2"
        }
      }

      storage_network {
        name                 = "Storage1Network"
        network_adapter_name = "StorageA"
        vlan_id              = "711"
      }

      storage_network {
        name                 = "Storage2Network"
        network_adapter_name = "StorageB"
        vlan_id              = "712"
      }
    }

    infrastructure_network {
      gateway      = "192.168.1.1"
      subnet_mask  = "255.255.255.0"
      dhcp_enabled = false
      dns_server = [
        "192.168.1.254"
      ]
      ip_pool {
        ending_address   = "192.168.1.65"
        starting_address = "192.168.1.55"
      }
    }

    optional_service {
      custom_location = "customlocation"
    }

    physical_node {
      ipv4_address = "192.168.1.12"
      name         = "AzSHOST1"
    }

    physical_node {
      ipv4_address = "192.168.1.13"
      name         = "AzSHOST2"
    }

    observability {
      streaming_data_client_enabled = true
      eu_location_enabled           = false
      episodic_data_upload_enabled  = true
    }

    security_setting {
      bitlocker_boot_volume_enabled   = true
      bitlocker_data_volume_enabled   = true
      credential_guard_enabled        = true
      drift_control_enabled           = true
      drtm_protection_enabled         = true
      hvci_protection_enabled         = true
      side_channel_mitigation_enabled = true
      smb_cluster_encryption_enabled  = false
      smb_signing_enabled             = true
      wdac_enabled                    = true
    }

    storage {
      configuration_mode = "Express"
    }
  }
}
