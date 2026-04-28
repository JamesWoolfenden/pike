resource "azurerm_oracle_exascale_database_storage_vault" "pike_gen" {
  name                              = "example-exascale-db-storage-vault"
  resource_group_name               = "pike"
  location                          = "pike"
  zones                             = ["1"]
  display_name                      = "example-exascale-db-storage-vault"
  description                       = "description"
  additional_flash_cache_percentage = 100
  high_capacity_database_storage {
    total_size_in_gb = 300
  }
  time_zone = "UTC"
}
