resource "google_bigtable_instance" "instance" {
  name = "bt-instance"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "europe-west2-a"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection = "false"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "bt-profile"

  // Requests will be routed to any of the 3 clusters.
  multi_cluster_routing_use_any = true

  ignore_warnings = true
}
