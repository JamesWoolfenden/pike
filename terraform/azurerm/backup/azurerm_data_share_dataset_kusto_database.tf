resource "azurerm_data_share_dataset_kusto_database" "pike_gen" {
  name              = "example-dskd"
  share_id          = "pike"
  kusto_database_id = "pike"
  depends_on = [
    azurerm_role_assignment.example,
  ]
}
