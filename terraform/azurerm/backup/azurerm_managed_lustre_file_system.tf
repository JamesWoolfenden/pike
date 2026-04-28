resource "azurerm_managed_lustre_file_system" "pike_gen" {
  name                   = "example-amlfs"
  resource_group_name    = "pike"
  location               = "pike"
  sku_name               = "AMLFS-Durable-Premium-250"
  subnet_id              = "pike"
  storage_capacity_in_tb = 8
  zones                  = ["2"]

  maintenance_window {
    day_of_week     = "Friday"
    time_of_day_utc = "22:00"
  }
}
