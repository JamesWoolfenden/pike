resource "azurerm_managed_disk_sas_token" "pike_gen" {
  managed_disk_id     = "pike"
  duration_in_seconds = 300
  access_level        = "Read"
}
