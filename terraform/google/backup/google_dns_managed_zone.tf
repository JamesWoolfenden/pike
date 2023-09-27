resource "google_dns_managed_zone" "pike" {
  name        = "pike"
  dns_name    = "freebeer.site."
  description = "Example DNS zone"
  labels = {
    foo = "bar"
  }
  private_visibility_config {
    gke_clusters {
      gke_cluster_name = google_container_cluster.cluster-1.id
    }
  }
}
