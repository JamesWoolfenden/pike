resource "azurerm_mssql_server_microsoft_support_auditing_policy" "pike_gen" {
  server_id             = "pike"
  blob_storage_endpoint = "pike"
}
