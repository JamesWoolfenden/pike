resource "azurerm_key_vault_managed_hardware_security_module" "pike_gen" {
  name                       = "exampleKVHsm"
  resource_group_name        = "pike"
  location                   = "pike"
  sku_name                   = "Standard_B1"
  purge_protection_enabled   = false
  soft_delete_retention_days = 90
  tenant_id                  = "pike"
  admin_object_ids           = ["pike"]

  tags = {
    Env = "Test"
  }
}
