resource "google_active_directory_peering" "pike" {
  provider           = google-beta
  domain_resource    = "pike"
  peering_id         = "ad-domain-peering"
  authorized_network = "projects/pangpt-ai/regions/us-central1/subnetworks/registry"

  labels = {
    foo = "bar"
  }
}
