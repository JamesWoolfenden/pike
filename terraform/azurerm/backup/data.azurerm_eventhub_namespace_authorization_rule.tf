data "azurerm_eventhub_namespace_authorization_rule" "pike_gen" {
  name                = "navi"
  resource_group_name = "example-resources"
  namespace_name      = "example-ns"
}
