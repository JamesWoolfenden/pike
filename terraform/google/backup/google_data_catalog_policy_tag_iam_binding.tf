data "google_iam_policy" "admin" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_data_catalog_policy_tag_iam_policy" "policy" {
  policy_tag  = google_data_catalog_policy_tag.pike.name
  policy_data = data.google_iam_policy.admin.policy_data
}
