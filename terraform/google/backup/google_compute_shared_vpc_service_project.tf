# Shared VPC Service Project - attaches a service project to a host project
# Note: For this test to work:
# - pike-477416 must already be configured as a Shared VPC host
# - pike-service-test must be a valid project you have access to

# First, manually enable the host (if not already done):
# gcloud compute shared-vpc enable pike-477416

resource "google_compute_shared_vpc_service_project" "pike" {
  host_project    = "pike-477416"
  service_project = "pike-477416" # Using same project for testing purposes only
}
