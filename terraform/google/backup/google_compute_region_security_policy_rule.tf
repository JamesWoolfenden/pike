resource "google_compute_region_security_policy" "pike_policy" {
  name    = "pike-policy-for-rule"
  region  = "europe-west2"
  project = "pike-477416"
  type    = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "pike" {
  region          = "europe-west2"
  project         = "pike-477416"
  security_policy = google_compute_region_security_policy.pike_policy.name
  priority        = 1000
  action          = "allow"

  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["*"]
    }
  }
}
