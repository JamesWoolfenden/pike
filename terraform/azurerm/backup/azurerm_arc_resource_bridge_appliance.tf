resource "azurerm_arc_resource_bridge_appliance" "pike_gen" {
  name                    = "example-appliance"
  location                = "pike"
  resource_group_name     = "pike"
  distro                  = "AKSEdge"
  infrastructure_provider = "VMWare"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    "hello" = "world"
  }
}
