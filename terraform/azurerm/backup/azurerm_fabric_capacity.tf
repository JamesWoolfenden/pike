resource "azurerm_fabric_capacity" "pike_gen" {
  name                = "exampleffc"
  resource_group_name = "pike"
  location            = "West Europe"

  administration_members = ["pike"]

  sku {
    name = "F32"
    tier = "Fabric"
  }

  tags = {
    environment = "test"
  }
}
