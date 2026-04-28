data "azurerm_eventhub" "pike_gen" {
  name                = "search-eventhub"
  resource_group_name = "search-service"
  namespace_name      = "search-eventhubns"
}
