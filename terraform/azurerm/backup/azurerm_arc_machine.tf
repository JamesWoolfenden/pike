resource "azurerm_arc_machine" "pike_gen" {
  name                = "example-arcmachine"
  resource_group_name = "pike"
  location            = "pike"
  kind                = "SCVMM"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    environment = "example"
  }
}
