resource "azurerm_key_vault" "example" {
  name                       = "pike"
  location                   = "uksouth"
  resource_group_name        = "pike"
  tenant_id                  = "8e7f742a-4215-44a0-881b-209124f286b1"
  sku_name                   = "premium"
  soft_delete_retention_days = 7
  access_policy {
    tenant_id = "8e7f742a-4215-44a0-881b-209124f286b1"
    object_id = "640b7b1e-da57-4518-a546-cbec8d9a9bce"

    key_permissions = [
      "Create",
      "Get",
      "Purge",
      "Recover"
    ]

    secret_permissions = [
      "Set",
    ]
  }
  purge_protection_enabled = true
  network_acls {
    default_action = "Deny"
    bypass         = "AzureServices"
  }
  tags = {
    pike = "permissions"
  }
}

output "key_vault" {
  value = azurerm_key_vault.example
}
