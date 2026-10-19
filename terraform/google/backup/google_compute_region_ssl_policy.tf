resource "google_compute_region_ssl_policy" "pike" {
  name    = "pike-test-ssl-policy"
  region  = "europe-west2"
  profile = "MODERN"
}