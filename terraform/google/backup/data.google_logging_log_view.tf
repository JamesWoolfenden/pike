data "google_logging_log_view" "pike" {
  parent   = "projects/pike-gcp"
  location = "global"
  bucket   = "_Default"
  name     = "_AllLogs"
}

output "google_logging_log_view" {
  value = data.google_logging_log_view.pike
}
