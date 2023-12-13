resource "google_storage_insights_report_config" "pike" {
  display_name = "Test Report Config"
  location     = "us-central1"
  frequency_options {
    frequency = "WEEKLY"
    start_date {
      day   = 15
      month = 3
      year  = 2050
    }
    end_date {
      day   = 15
      month = 4
      year  = 2050
    }
  }
  csv_options {
    record_separator = "\n"
    delimiter        = ","
    header_required  = false
  }
}
