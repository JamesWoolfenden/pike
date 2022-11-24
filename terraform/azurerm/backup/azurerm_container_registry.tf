resource "azurerm_container_registry" "pike" {
  name                          = "pikeRepo"
  resource_group_name           = "pike"
  location                      = "uk south"
  admin_enabled                 = false
  anonymous_pull_enabled        = false
  sku                           = "Premium"
  public_network_access_enabled = true

  georeplications {
    location                = "East US"
    zone_redundancy_enabled = true
    tags = {
      pike = "permissions"
    }
  }

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.example.id]
  }

  encryption {
    enabled            = true
    identity_client_id = azurerm_user_assigned_identity.example.client_id
    key_vault_key_id   = data.azurerm_key_vault_key.pike.id
  }

  tags = {
    pike    = "permissions"
    another = "here"
  }
}

resource "azurerm_user_assigned_identity" "example" {
  resource_group_name = "pike"
  location            = "uk south"

  name = "registry-uai"
}

data "azurerm_key_vault_key" "pike" {
  name         = "pike"
  key_vault_id = data.azurerm_key_vault.pike.id
}

data "azurerm_key_vault" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
