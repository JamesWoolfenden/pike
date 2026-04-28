resource "azurerm_network_profile" "pike_gen" {
  name                = "examplenetprofile"
  location            = "pike"
  resource_group_name = "pike"

  container_network_interface {
    name = "examplecnic"

    ip_configuration {
      name      = "exampleipconfig"
      subnet_id = "pike"
    }
  }
}
