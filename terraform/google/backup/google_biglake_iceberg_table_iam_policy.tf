data "google_iam_policy" "google_biglake_iceberg_table" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_biglake_iceberg_table_iam_policy" "pike" {
  name      = google_biglake_iceberg_table.pike.name
  catalog   = google_biglake_iceberg_catalog.pike.name
  namespace = google_biglake_iceberg_namespace.pike.namespace_id

  policy_data = data.google_iam_policy.google_biglake_iceberg_table.policy_data
}
