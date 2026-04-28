resource "azurerm_mongo_cluster" "pike_gen" {
  name                = "example-mc"
  resource_group_name = "pike"
  location            = "pike"

  administrator_username = "adminTerraform"
  compute_tier           = "Free"
  high_availability_mode = "Disabled"
  shard_count            = "1"
  storage_size_in_gb     = "32"
  version                = "8.0"
}
