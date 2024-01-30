resource "google_spanner_database_iam_binding" "database" {
  instance = google_spanner_instance.example.name
  database = google_spanner_database.database.name
  role     = "roles/editor"

  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
