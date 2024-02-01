resource "google_bigquery_connection_iam_binding" "pike" {
  connection_id = google_bigquery_connection.pike.connection_id

  role = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
