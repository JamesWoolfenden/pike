data "google_vmwareengine_subnet" "pike" {
  name   = "service-1"
  parent = "projects/pike-gcp/locations/us-central1/privateClouds/my-cloud"
}
