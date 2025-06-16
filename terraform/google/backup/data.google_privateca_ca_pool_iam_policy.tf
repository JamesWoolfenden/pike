data "google_privateca_ca_pool_iam_policy" "pike" {
  ca_pool = google_privateca_ca_pool.default.id
}

output "google_privateca_ca_pool_iam_policy" {
  value = data.google_privateca_ca_pool_iam_policy.pike
}

resource "google_privateca_ca_pool" "default" {
  name     = "my-pool"
  location = "us-central1"
  tier     = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = true
  }
  labels = {
    foo = "bar"
  }
}
