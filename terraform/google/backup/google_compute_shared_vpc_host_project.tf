# Note: Shared VPC requires actual projects with proper IAM setup
# This is a minimal test to discover the required permissions

resource "google_compute_shared_vpc_host_project" "pike" {
  project = "pike-477416"
}

# Service project attachment (requires a separate project)
# Commenting out as it needs another project
# resource "google_compute_shared_vpc_service_project" "pike" {
#   host_project    = google_compute_shared_vpc_host_project.pike.project
#   service_project = "pike-service-project"
# }
