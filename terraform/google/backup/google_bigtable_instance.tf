resource "google_bigtable_instance" "pike" {
  name = "tf-instance"

  cluster {
    cluster_id   = "tf-instance-cluster"
    num_nodes    = 1
    storage_type = "HDD"
    zone         = "europe-west2-a"
  }

  labels = {
    my-label = "prod-label"
  }

  deletion_protection = false
}
