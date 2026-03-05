resource "google_logging_project_sink" "pike" {
  name        = "pike"
  destination = "storage.googleapis.com/terraform-pike-bucket-tfstate"

  unique_writer_identity = true
}
