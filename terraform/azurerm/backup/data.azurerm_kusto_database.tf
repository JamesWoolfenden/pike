data "azurerm_kusto_database" "pike_gen" {
  name                = "my-kusto-database"
  resource_group_name = "test_resource_group"
  cluster_name        = "test_cluster"
}
