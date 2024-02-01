data "google_iam_policy" "admin2" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_cloud_run_v2_job_iam_binding" "policy" {
  name = google_cloud_run_v2_job.default.name
  role = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}

resource "google_cloud_run_v2_job" "default" {
  name     = "cloudrun-job"
  location = "europe-west2"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}
