resource "azurerm_application_security_group" "pike_gen" {
  name                = "tf-appsecuritygroup"
  location            = "pike"
  resource_group_name = "pike"

  tags = {
    Hello = "World"
  }
}
