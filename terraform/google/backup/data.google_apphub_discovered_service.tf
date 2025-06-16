data "google_apphub_discovered_service" "pike" {
  location    = "us-central1"
  service_uri = "/compute.googleapis.com/projects/1/regions/us-east1/instanceGroups/id1"
}

output "google_apphub_discovered_service" {
  value = data.google_apphub_discovered_service.pike
}
