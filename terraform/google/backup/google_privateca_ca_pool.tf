resource "google_privateca_ca_pool" "pike" {
  name     = "my-pool2"
  location = "us-central1"
  tier     = "DEVOPS"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = false
  }
  labels = {
    foo = "bar"
  }
}
