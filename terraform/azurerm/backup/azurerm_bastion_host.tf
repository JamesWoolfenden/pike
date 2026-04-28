resource "azurerm_bastion_host" "pike_gen" {
  name                = "examplebastion"
  location            = "pike"
  resource_group_name = "pike"

  ip_configuration {
    name                 = "configuration"
    subnet_id            = "pike"
    public_ip_address_id = "pike"
  }
}
