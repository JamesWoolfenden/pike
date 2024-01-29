resource "google_firebase_project" "pike" {
  provider = google-beta
  project  = var.name
}

variable "name" {
  default = "pike-gcp"
}
