resource "google_spanner_database_iam_member" "database" {
  instance = google_spanner_instance.example.name
  database = google_spanner_database.database.name
  role     = "roles/editor"
  member   = "user:crwoolfenden@gmail.com"
}
