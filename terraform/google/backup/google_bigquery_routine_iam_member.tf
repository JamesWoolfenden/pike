resource "google_bigquery_routine_iam_member" "pike" {
  dataset_id = google_bigquery_dataset.pike.dataset_id
  routine_id = google_bigquery_routine.pike.routine_id

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
