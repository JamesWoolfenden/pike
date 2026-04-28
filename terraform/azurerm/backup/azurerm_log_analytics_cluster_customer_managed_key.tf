resource "azurerm_log_analytics_cluster_customer_managed_key" "pike_gen" {
  log_analytics_cluster_id = "pike"
  key_vault_key_id         = "pike"
}
