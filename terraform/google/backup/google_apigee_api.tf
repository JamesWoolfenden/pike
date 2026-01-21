resource "google_apigee_api" "pike" {
  name          = "pike"
  org_id        = google_apigee_organization.org.id
  config_bundle = data.archive_file.bundle.output_path
}

data "archive_file" "bundle" {
  type             = "zip"
  source_dir       = "${path.module}/role"
  output_path      = "${path.module}/bundle.zip"
  output_file_mode = "0644"
}
