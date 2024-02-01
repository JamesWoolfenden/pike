resource "google_bigquery_routine" "pike" {
  dataset_id      = google_bigquery_dataset.default.dataset_id
  routine_id      = "routine_id"
  routine_type    = "PROCEDURE"
  language        = "SQL"
  definition_body = "CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);"

}
