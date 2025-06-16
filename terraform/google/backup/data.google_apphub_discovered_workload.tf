data "google_apphub_discovered_workload" "pike" {
  location     = "us-central1"
  workload_uri = "/compute.googleapis.com/projects/1/regions/us-east1/instanceGroups/id1"
}

output "google_apphub_discovered_workload" {
  value = data.google_apphub_discovered_workload.pike
}
