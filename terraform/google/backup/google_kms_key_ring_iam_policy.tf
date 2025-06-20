data "google_iam_policy" "admin" {
  binding {
    role = "roles/editor"

    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_kms_key_ring_iam_policy" "key_ring" {
  key_ring_id = google_kms_key_ring.keyring.id
  policy_data = data.google_iam_policy.admin.policy_data
}
