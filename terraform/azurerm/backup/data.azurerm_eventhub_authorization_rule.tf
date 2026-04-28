data "azurerm_eventhub_authorization_rule" "pike_gen" {
  name                = "test"
  namespace_name      = "pike"
  eventhub_name       = "pike"
  resource_group_name = "pike"
}
