resource "azurerm_oracle_autonomous_database" "pike_gen" {
  name                             = "example"
  resource_group_name              = "example"
  location                         = "West Europe"
  subnet_id                        = "example"
  display_name                     = "example"
  db_workload                      = "example"
  mtls_connection_required         = false
  backup_retention_period_in_days  = 42
  compute_model                    = "example"
  data_storage_size_in_gbs         = 42
  auto_scaling_for_storage_enabled = false
  virtual_network_id               = "example"
  auto_scaling_enabled             = "example"
  character_set                    = "example"
  compute_count                    = 1.23456
  national_character_set           = "example"
  license_model                    = false
  db_version                       = "example"
}
