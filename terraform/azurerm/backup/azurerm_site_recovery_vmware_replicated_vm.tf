resource "azurerm_site_recovery_vmware_replicated_vm" "pike_gen" {
  name                                       = "example-vmware-vm"
  recovery_vault_id                          = "pike"
  source_vm_name                             = "example-vm"
  appliance_name                             = "example-appliance"
  recovery_replication_policy_id             = "pike"
  physical_server_credential_name            = "example-creds"
  license_type                               = "NotSpecified"
  target_boot_diagnostics_storage_account_id = "pike"
  target_vm_name                             = "example_replicated_vm"
  target_resource_group_id                   = "pike"
  default_log_storage_account_id             = "pike"
  default_recovery_disk_type                 = "Standard_LRS"
  target_network_id                          = "pike"

  network_interface {
    source_mac_address = "00:00:00:00:00:00"
    target_subnet_name = "pike"
    is_primary         = true
  }
}
