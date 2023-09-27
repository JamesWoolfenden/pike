
data "google_iam_policy" "admin" {
  binding {
    role = "roles/bigtable.user"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_bigtable_instance_iam_policy" "editor" {
  project     = "pike-gcp"
  instance    = google_bigtable_instance.instance.name
  policy_data = data.google_iam_policy.admin.policy_data
}
