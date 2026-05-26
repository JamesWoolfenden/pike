data "google_artifact_registry_file" "pike" {
  location      = "us-central1"
  repository_id = "pike"
  file_id       = "pike"
  output_path   = "/tmp/pike"
}

output "google_artifact_registry_file" {
  value = data.google_artifact_registry_file.pike
}
