data "google_logging_sink" "pike" {
  id = google_logging_project_sink.pike.id
}

output "google_logging_sink" {
  value = data.google_logging_sink.pike
}
