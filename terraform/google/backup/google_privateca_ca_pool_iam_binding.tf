resource "google_privateca_ca_pool_iam_binding" "binding" {
  ca_pool = google_privateca_ca_pool.pike.id
  role    = "roles/privateca.certificateManager"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
