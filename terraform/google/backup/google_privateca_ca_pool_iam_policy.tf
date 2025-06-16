data "google_iam_policy" "admin2" {
  binding {
    role = "roles/privateca.certificateManager"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_privateca_ca_pool_iam_policy" "policy" {
  ca_pool     = google_privateca_ca_pool.pike.id
  policy_data = data.google_iam_policy.admin2.policy_data
}
