resource "azurerm_databricks_access_connector" "pike_gen" {
  name                = "example-resource"
  resource_group_name = "pike"
  location            = "pike"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }
}
