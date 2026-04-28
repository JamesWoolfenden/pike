resource "azurerm_kusto_database_principal_assignment" "pike_gen" {
  name                = "KustoPrincipalAssignment"
  resource_group_name = "pike"
  cluster_name        = "pike"
  database_name       = "pike"

  tenant_id      = "pike"
  principal_id   = "pike"
  principal_type = "App"
  role           = "Viewer"
}
