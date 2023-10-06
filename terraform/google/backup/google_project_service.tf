resource "google_project_service" "project" {
  //project = "pike-gcp"
  service = "bigquerymigration.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}
