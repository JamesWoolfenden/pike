data "azurerm_kusto_cluster" "pike_gen" {
  name                = "kustocluster"
  resource_group_name = "test_resource_group"
}
