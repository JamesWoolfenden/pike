data "google_iam_policy" "admin2" {
  binding {
    role = "roles/editor"

    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_spanner_database_iam_policy" "database" {
  instance    = google_spanner_instance.example.name
  database    = google_spanner_database.database.name
  policy_data = data.google_iam_policy.admin2.policy_data
}
