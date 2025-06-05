resource "google_composer_environment" "pike" {
  name   = "pike-test"
  region = "europe-west2"

  labels = {
    "pike" = "permissions"
  }

  config {

    software_config {
      image_version = "composer-3-airflow-2"
    }

    environment_size = "ENVIRONMENT_SIZE_SMALL"

    node_config {
      service_account = google_service_account.test.name
    }
  }
}


resource "google_service_account" "test" {
  account_id   = "composer-env-account"
  display_name = "Test Service Account for Composer Environment"
}

resource "google_project_iam_member" "composer-worker" {
  project = "pike-412922"
  role    = "roles/composer.worker"
  member  = "serviceAccount:${google_service_account.test.email}"
}
