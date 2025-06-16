resource "google_privateca_ca_pool_iam_member" "member" {
  ca_pool = google_privateca_ca_pool.pike.id
  role    = "roles/privateca.certificateManager"
  member  = "user:james.woolfenden@gmail.com"
}
