resource "azurerm_digital_twins_instance" "pike_gen" {
  name                = "example-DT"
  resource_group_name = "pike"
  location            = "pike"

  tags = {
    foo = "bar"
  }
}
