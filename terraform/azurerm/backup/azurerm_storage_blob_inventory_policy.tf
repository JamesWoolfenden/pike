resource "azurerm_storage_blob_inventory_policy" "pike_gen" {
  storage_account_id = "pike"
  rules {
    name                   = "rule1"
    storage_container_name = "pike"
    format                 = "Csv"
    schedule               = "Daily"
    scope                  = "Container"
    schema_fields = [
      "Name",
      "Last-Modified",
    ]
  }
}
