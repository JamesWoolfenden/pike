resource "google_project_service" "project" {
  project = "pike-477416"
  service = "iam.googleapis.com"
}

resource "google_project_service" "cloudresources" {
  project = "pike-477416"
  service = "cloudresourcemanager.googleapis.com"
}
