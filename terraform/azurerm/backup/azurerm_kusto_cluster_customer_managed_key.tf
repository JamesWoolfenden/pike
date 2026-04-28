resource "azurerm_kusto_cluster_customer_managed_key" "pike_gen" {
  cluster_id   = "pike"
  key_vault_id = "pike"
  key_name     = "pike"
  key_version  = "pike"
}
