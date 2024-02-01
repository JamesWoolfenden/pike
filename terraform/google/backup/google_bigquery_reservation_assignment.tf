resource "google_bigquery_reservation_assignment" "pike" {
  assignee    = "projects/pike-gcp"
  job_type    = "PIPELINE"
  reservation = google_bigquery_reservation.pike.id
}
