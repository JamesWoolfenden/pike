resource "azurerm_mssql_managed_instance_start_stop_schedule" "pike_gen" {
  managed_instance_id = "pike"
  timezone_id         = "Central European Standard Time"

  schedule {
    start_day  = "Monday"
    start_time = "08:00"
    stop_day   = "Monday"
    stop_time  = "11:00"
  }

  schedule {
    start_day  = "Tuesday"
    start_time = "12:00"
    stop_day   = "Tuesday"
    stop_time  = "18:00"
  }
}
