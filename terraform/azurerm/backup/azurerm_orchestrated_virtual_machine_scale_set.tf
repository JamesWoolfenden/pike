resource "azurerm_orchestrated_virtual_machine_scale_set" "pike_gen" {
  name                = "example-VMSS"
  location            = "pike"
  resource_group_name = "pike"

  platform_fault_domain_count = 1

  zones = ["1"]
}
