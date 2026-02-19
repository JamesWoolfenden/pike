resource "google_compute_firewall_policy" "pike" {
  parent      = "organizations/123456789"
  short_name  = "pike-test-policy"
  description = "Test firewall policy for pike"
}
