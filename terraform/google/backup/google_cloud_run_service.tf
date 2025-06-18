resource "google_cloud_run_service" "pike" {
  name     = "cloudrun-srv"
  location = "europe-west2"

  template {

    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

}
