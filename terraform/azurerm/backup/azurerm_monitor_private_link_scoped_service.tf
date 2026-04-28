resource "azurerm_monitor_private_link_scoped_service" "pike_gen" {
  name                = "example-amplsservice"
  resource_group_name = "pike"
  scope_name          = "pike"
  linked_resource_id  = "pike"
}
