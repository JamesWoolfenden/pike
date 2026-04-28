resource "azurerm_data_share_account" "pike_gen" {
  name                = "example-dsa"
  location            = "pike"
  resource_group_name = "pike"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    foo = "bar"
  }
}
