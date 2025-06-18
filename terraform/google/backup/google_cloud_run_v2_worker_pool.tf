resource "google_cloud_run_v2_worker_pool" "pike" {
  name                = "cloudrun-worker-pool"
  location            = "europe-west2"
  deletion_protection = false
  launch_stage        = "ALPHA"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/worker-pool"
    }
  }
}
