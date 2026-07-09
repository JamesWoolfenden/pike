data "google_iam_policy" "google_bigquery_routine" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_routine_iam_policy" "pike" {
  dataset_id = google_bigquery_dataset.pike.dataset_id
  routine_id = google_bigquery_routine.pike.routine_id

  policy_data = data.google_iam_policy.google_bigquery_routine.policy_data
}
