resource "google_bigquery_routine" "pike" {
  dataset_id      = google_bigquery_dataset.pike.dataset_id
  routine_id      = "pike"
  routine_type    = "SCALAR_FUNCTION"
  language        = "SQL"
  definition_body = "1 + value"
  arguments {
    name      = "value"
    data_type = "{\"typeKind\" :  \"INT64\"}"
  }
  return_type = "{\"typeKind\" :  \"INT64\"}"
}
