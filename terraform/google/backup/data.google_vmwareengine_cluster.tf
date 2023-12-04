data "google_vmwareengine_cluster" "pike" {
  provider = google-beta
  parent   = "pikeparent"
  name     = "pike"
}
