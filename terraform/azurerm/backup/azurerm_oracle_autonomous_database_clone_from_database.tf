resource "azurerm_oracle_autonomous_database_clone_from_database" "pike_gen" {
  name                          = "example"
  resource_group_name           = "pike"
  location                      = "pike"
  source_autonomous_database_id = "pike"

  clone_type = "Full"

  backup_retention_period_in_days  = 7
  character_set                    = "AL32UTF8"
  compute_count                    = 2.0
  compute_model                    = "ECPU"
  data_storage_size_in_tb          = 1
  database_version                 = "19c"
  database_workload                = "OLTP"
  display_name                     = "ExampleClone"
  license_model                    = "LicenseIncluded"
  auto_scaling_enabled             = false
  auto_scaling_for_storage_enabled = true
  mtls_connection_required         = true
  national_character_set           = "AL16UTF16"
  allowed_ip_addresses             = []
}
