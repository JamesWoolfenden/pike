data "google_vmwareengine_vcenter_credentials" "pike" {
  parent = "projects/pike-gcp/locations/us-central1/privateClouds/my-cloud"
}
