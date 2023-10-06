data "google_project_service" "pike" {
  service = "gkebackup.googleapis.com"
}

resource "google_project_service" "project" {
  //project = "pike-gcp"
  service = "gkebackup.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}
