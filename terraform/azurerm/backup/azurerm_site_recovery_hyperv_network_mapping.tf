resource "azurerm_site_recovery_hyperv_network_mapping" "pike_gen" {
  name                                              = "recovery-network-mapping"
  recovery_vault_id                                 = "pike"
  source_system_center_virtual_machine_manager_name = "my-vmm-server"
  source_network_name                               = "my-vmm-network"
  target_network_id                                 = "pike"
}
