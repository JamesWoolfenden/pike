resource "google_logging_project_bucket_config" "pike" {
  project        = "pike-477416"
  location       = "global"
  bucket_id      = "_Default"
  retention_days = 30
  description    = "Pike test bucket config - updated"
}
