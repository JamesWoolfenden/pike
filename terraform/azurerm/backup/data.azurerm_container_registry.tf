data "azurerm_container_registry" "pike_gen" {
  name                = "testacr"
  resource_group_name = "test"
}
