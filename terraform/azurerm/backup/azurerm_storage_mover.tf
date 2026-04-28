resource "azurerm_storage_mover" "pike_gen" {
  name                = "example-ssm"
  resource_group_name = "pike"
  location            = "West Europe"
  description         = "Example Storage Mover Description"
  tags = {
    key = "value"
  }
}
