resource "azurerm_video_indexer_account" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "West Europe"

  storage {
    storage_account_id = "pike"
  }

  identity {
    type = "SystemAssigned"
  }
}
