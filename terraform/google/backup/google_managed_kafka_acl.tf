resource "google_managed_kafka_cluster" "cluster" {
  cluster_id = "my-cluster"
  location   = "us-central1"
  capacity_config {
    vcpu_count   = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
}

resource "google_managed_kafka_acl" "example" {
  acl_id   = "topic/mytopic"
  cluster  = google_managed_kafka_cluster.cluster.cluster_id
  location = "us-central1"
  acl_entries {
    principal       = "User:admin@pike-412922.iam.gserviceaccount.com"
    permission_type = "ALLOW"
    operation       = "ALL"
    host            = "*"
  }
  acl_entries {
    principal       = "User:producer-client@pike-412922.iam.gserviceaccount.com"
    permission_type = "ALLOW"
    operation       = "WRITE"
    host            = "*"
  }

  depends_on = [
    google_managed_kafka_cluster.cluster
  ]
}

data "google_project" "project" {
}
