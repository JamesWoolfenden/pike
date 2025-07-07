

resource "google_dataform_repository_release_config" "release" {
  provider = google-beta

  project    = google_dataform_repository.pike.project
  region     = google_dataform_repository.pike.region
  repository = google_dataform_repository.pike.name

  name          = "my_release"
  git_commitish = "main"
  cron_schedule = "0 7 * * *"
  time_zone     = "America/New_York"

  code_compilation_config {
    default_database = "gcp-example-project"
    default_schema   = "example-dataset"
    default_location = "us-central1"
    assertion_schema = "example-assertion-dataset"
    database_suffix  = ""
    schema_suffix    = ""
    table_prefix     = ""
    vars = {
      var1 = "value"
    }
  }
}
