resource "azurerm_databricks_workspace" "pike_gen" {
  name                = "databricks-test"
  resource_group_name = "pike"
  location            = "pike"
  sku                 = "standard"

  tags = {
    Environment = "Production"
  }
}
