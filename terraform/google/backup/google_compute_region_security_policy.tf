resource "google_compute_region_security_policy" "pike" {
  name        = "pike-region-security-policy"
  region      = "europe-west2"
  project     = "pike-477416"
  type        = "CLOUD_ARMOR"
  description = "Test security policy for Pike"
}
