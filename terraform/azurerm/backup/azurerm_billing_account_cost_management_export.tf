resource "azurerm_billing_account_cost_management_export" "pike_gen" {
  name                         = "example"
  billing_account_id           = "example"
  recurrence_type              = "Monthly"
  recurrence_period_start_date = "2020-08-18T00:00:00Z"
  recurrence_period_end_date   = "2020-09-18T00:00:00Z"
  file_format                  = "Csv"

  export_data_storage_location {
    container_id     = "pike"
    root_folder_path = "/root/updated"
  }

  export_data_options {
    type       = "Usage"
    time_frame = "WeekToDate"
  }
}
