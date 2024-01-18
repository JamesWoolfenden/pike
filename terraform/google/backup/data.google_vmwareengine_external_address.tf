data "google_vmwareengine_external_address" "pike" {
  name   = "my-external-address"
  parent = "projects/pike-gcp/locations/us-central1/privateClouds/my-cloud"
}
