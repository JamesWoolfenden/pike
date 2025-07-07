
resource "google_service_account" "dataform_sa" {
  provider     = google-beta
  account_id   = "dataform-sa"
  display_name = "Dataform Service Account"
}

resource "google_dataform_repository_workflow_config" "workflow" {
  provider = google-beta

  project        = google_dataform_repository.pike.project
  region         = google_dataform_repository.pike.region
  repository     = google_dataform_repository.pike.name
  name           = "my_workflow"
  release_config = google_dataform_repository_release_config.release.id

  invocation_config {
    included_targets {
      database = "gcp-example-project"
      schema   = "example-dataset"
      name     = "target_1"
    }
    included_targets {
      database = "gcp-example-project"
      schema   = "example-dataset"
      name     = "target_2"
    }
    included_tags                            = ["tag_1"]
    transitive_dependencies_included         = true
    transitive_dependents_included           = true
    fully_refresh_incremental_tables_enabled = false
    service_account                          = google_service_account.dataform_sa.email
  }

  cron_schedule = "0 7 * * *"
  time_zone     = "America/New_York"
}
