resource "azurerm_mssql_managed_database" "pike_gen" {
  name                = "example"
  managed_instance_id = "pike"

  # prevent the possibility of accidental data loss
  lifecycle {
    prevent_destroy = true
  }
}
