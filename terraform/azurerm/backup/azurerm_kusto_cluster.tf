resource "azurerm_kusto_cluster" "pike_gen" {
  name                = "example"
  location            = "pike"
  resource_group_name = "pike"

  sku {
    name     = "Standard_D13_v2"
    capacity = 2
  }

  tags = {
    Environment = "Production"
  }
}
