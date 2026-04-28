data "azurerm_eventhub_cluster" "pike_gen" {
  name                = "search-eventhub"
  resource_group_name = "search-service"
}
