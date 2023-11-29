data "google_monitoring_istio_canonical_service" "pike" {
  provider                    = google-beta
  mesh_uid                    = "proj-573164786102"
  canonical_service_namespace = "istio-system"
  canonical_service           = "prometheus"
}
