data "azurerm_dedicated_host" "pike_gen" {
  name                      = "example-host"
  dedicated_host_group_name = "example-host-group"
  resource_group_name       = "example-resources"
}
