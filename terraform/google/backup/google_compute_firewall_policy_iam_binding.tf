resource "google_compute_firewall_policy_iam_binding" "pike" {
  name = google_compute_firewall_policy.pike.name

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
