resource "google_secret_manager_secret" "db_password" {
  secret_id = "db-password"
  replication {
    auto {}
  }
}

resource "google_cloud_run_v2_service" "app" {
  name     = "my-app"
  location = "us-central1"

  template {
    containers {
      image = "gcr.io/my-project/my-app:latest"

      env {
        name = "DB_PASSWORD"
        value_source {
          secret_key_ref {
            secret = google_secret_manager_secret.db_password.secret_id
          }
        }
      }

      env {
        name  = "REGULAR_VAR"
        value = "not-a-secret"
      }
    }

    cloud_sql_instances {
      instance = "my-project:us-central1:my-db"
    }
  }
}
