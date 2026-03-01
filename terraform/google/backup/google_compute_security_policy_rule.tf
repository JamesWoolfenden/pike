# Security Policy Rule - Cloud Armor rule for DDoS protection
resource "google_compute_security_policy" "pike_policy" {
  name    = "pike-security-policy"
  project = "pike-477416"
}

resource "google_compute_security_policy_rule" "pike" {
  security_policy = google_compute_security_policy.pike_policy.name
  priority        = 1000
  project         = "pike-477416"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.0.2.0/24"]
    }
  }

  action = "deny(403)"
}
