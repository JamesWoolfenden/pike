data "google_monitoring_mesh_istio_service" "pike" {
  provider          = google-beta
  service_namespace = "pike"
  mesh_uid          = "proj-573164786102"
  service_name      = "pike"
}
