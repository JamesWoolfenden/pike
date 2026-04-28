resource "azurerm_databricks_workspace_root_dbfs_customer_managed_key" "pike_gen" {
  depends_on = ["pike"]

  workspace_id     = "pike"
  key_vault_key_id = "pike"
}
