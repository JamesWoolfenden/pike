data "google_monitoring_cluster_istio_service" "pike" {
  cluster_name      = "pike"
  location          = "us-central1"
  service_name      = "pike"
  service_namespace = "pike"
}
