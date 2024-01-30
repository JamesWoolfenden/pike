resource "google_spanner_instance_iam_member" "instance" {
  instance = google_spanner_instance.example.name
  role     = "roles/spanner.databaseAdmin"
  member   = "user:crwoolfenden@gmail.com"
}
