resource "azurerm_site_recovery_replicated_vm" "pike_gen" {
  name                                      = "vm-replication"
  resource_group_name                       = "pike"
  recovery_vault_name                       = "pike"
  source_recovery_fabric_name               = "pike"
  source_vm_id                              = "pike"
  recovery_replication_policy_id            = "pike"
  source_recovery_protection_container_name = "pike"

  target_resource_group_id                = "pike"
  target_recovery_fabric_id               = "pike"
  target_recovery_protection_container_id = "pike"

  managed_disk {
    disk_id                    = "pike"
    staging_storage_account_id = "pike"
    target_resource_group_id   = "pike"
    target_disk_type           = "Premium_LRS"
    target_replica_disk_type   = "Premium_LRS"
  }

  network_interface {
    source_network_interface_id   = "pike"
    target_subnet_name            = "pike"
    recovery_public_ip_address_id = "pike"
  }

  depends_on = [
    azurerm_site_recovery_protection_container_mapping.container-mapping,
    azurerm_site_recovery_network_mapping.network-mapping,
  ]
}
