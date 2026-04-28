data "azurerm_iothub_shared_access_policy" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  iothub_name         = "pike"
}
