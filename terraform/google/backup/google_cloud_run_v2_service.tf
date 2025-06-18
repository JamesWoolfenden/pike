resource "google_cloud_run_v2_service" "pike" {
  name                = "pike"
  location            = "europe-west2"
  deletion_protection = false
  ingress             = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = "super-544@pike-412922.iam.gserviceaccount.com"
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}
