data "azurerm_eventhub_namespace" "pike_gen" {
  name                = "search-eventhubns"
  resource_group_name = "search-service"
}
