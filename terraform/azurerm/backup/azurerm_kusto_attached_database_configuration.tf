resource "azurerm_kusto_attached_database_configuration" "pike_gen" {
  name                = "configuration1"
  resource_group_name = "pike"
  location            = "pike"
  cluster_name        = "pike"
  cluster_id          = "pike"
  database_name       = "pike"

  sharing {
    external_tables_to_exclude    = ["ExternalTable2"]
    external_tables_to_include    = ["ExternalTable1"]
    functions_to_exclude          = ["Function2"]
    functions_to_include          = ["Function1"]
    materialized_views_to_exclude = ["MaterializedViewTable2"]
    materialized_views_to_include = ["MaterializedViewTable1"]
    tables_to_exclude             = ["Table2"]
    tables_to_include             = ["Table1"]
  }
}
