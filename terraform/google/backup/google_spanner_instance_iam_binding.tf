resource "google_spanner_instance_iam_binding" "instance" {
  instance = google_spanner_instance.example.name
  role     = "roles/spanner.databaseAdmin"

  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
