resource "azurerm_stack_hci_cluster" "pike_gen" {
  name                = "example-cluster"
  resource_group_name = "pike"
  location            = "pike"
  client_id           = "pike"
  tenant_id           = "pike"
  identity {
    type = "SystemAssigned"
  }
}
