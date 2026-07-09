resource "google_biglake_iceberg_table_iam_member" "pike" {
  name      = google_biglake_iceberg_table.pike.name
  catalog   = google_biglake_iceberg_catalog.pike.name
  namespace = google_biglake_iceberg_namespace.pike.namespace_id

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
