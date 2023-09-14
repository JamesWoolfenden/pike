resource "google_bigquery_job" "job" {
  job_id = "job_query${random_string.random.id}"

  labels = {
    pike            = "permissions"
    "example-label" = "example-value"
  }

  query {
    query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]"

    destination_table {
      project_id = google_bigquery_table.pike.project
      dataset_id = google_bigquery_table.pike.dataset_id
      table_id   = google_bigquery_table.pike.table_id
    }

    allow_large_results = true
    flatten_results     = true

    script_options {
      key_result_statement = "LAST"
    }
  }
}

resource "random_string" "random" {
  length  = 8
  special = false
}
