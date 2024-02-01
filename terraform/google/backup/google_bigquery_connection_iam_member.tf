resource "google_bigquery_connection_iam_member" "pike" {
  connection_id = google_bigquery_connection.pike.connection_id
  role          = "roles/viewer"
  member        = "user:crwoolfenden@gmail.com"
}
