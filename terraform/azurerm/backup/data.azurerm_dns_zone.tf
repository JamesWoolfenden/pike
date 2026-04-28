data "azurerm_dns_zone" "pike_gen" {
  name                = "search-eventhubns"
  resource_group_name = "search-service"
}
