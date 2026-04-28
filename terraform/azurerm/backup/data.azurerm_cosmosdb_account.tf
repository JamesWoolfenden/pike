data "azurerm_cosmosdb_account" "pike_gen" {
  name                = "tfex-cosmosdb-account"
  resource_group_name = "tfex-cosmosdb-account-rg"
}
