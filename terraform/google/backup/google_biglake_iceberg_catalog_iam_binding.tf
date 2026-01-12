resource "google_biglake_iceberg_catalog_iam_binding" "pike" {
  project = google_biglake_iceberg_catalog.pike.project
  name    = google_biglake_iceberg_catalog.pike.name
  role    = "roles/biglake.editor"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
