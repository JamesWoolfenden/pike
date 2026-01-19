resource "google_healthcare_pipeline_job" "pike" {
  dataset  = "data-1234546"
  location = "us-central1"
  name     = "pike"
}
