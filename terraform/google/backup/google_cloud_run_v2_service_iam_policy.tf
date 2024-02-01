data "google_iam_policy" "admin" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_cloud_run_v2_service_iam_policy" "pike" {

  name        = google_cloud_run_v2_service.pike.name
  policy_data = data.google_iam_policy.admin.policy_data
}

resource "google_cloud_run_v2_service" "pike" {
  name     = "pike"
  ingress  = "INGRESS_TRAFFIC_ALL"
  location = "europe-west2"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}
