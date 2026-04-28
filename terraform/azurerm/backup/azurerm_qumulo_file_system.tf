resource "azurerm_qumulo_file_system" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  storage_sku         = "Standard"
  subnet_id           = "pike"
  zone                = "1"
}
