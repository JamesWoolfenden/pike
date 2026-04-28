resource "azurerm_data_share_dataset_kusto_cluster" "pike_gen" {
  name             = "example-dskc"
  share_id         = "pike"
  kusto_cluster_id = "pike"
  depends_on = [
    azurerm_role_assignment.example,
  ]
}
