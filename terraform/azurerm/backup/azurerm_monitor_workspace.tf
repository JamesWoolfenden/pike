resource "azurerm_monitor_workspace" "pike_gen" {
  name                = "example-mamw"
  resource_group_name = "pike"
  location            = "West Europe"
  tags = {
    key = "value"
  }
}
