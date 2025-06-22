resource "google_data_catalog_taxonomy_iam_binding" "binding" {
  taxonomy = google_data_catalog_taxonomy.pike.name
  role     = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
