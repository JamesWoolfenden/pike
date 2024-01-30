data "google_iam_policy" "admin" {
  binding {
    role = "roles/editor"

    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_spanner_instance_iam_policy" "instance" {
  instance    = google_spanner_instance.example.name
  policy_data = data.google_iam_policy.admin.policy_data
}
