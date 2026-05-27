data "google_iam_policy" "admin3" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_data_catalog_entry_group_iam_policy" "policy" {
  entry_group = google_data_catalog_entry_group.pike.name
  policy_data = data.google_iam_policy.admin.policy_data
}
