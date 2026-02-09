resource "google_clouddeploy_delivery_pipeline" "pike" {
  name     = "cd-pipeline"
  location = "us-central1"
  serial_pipeline {
    stages {
      target_id = "test"
      profiles  = []
    }
  }
}
