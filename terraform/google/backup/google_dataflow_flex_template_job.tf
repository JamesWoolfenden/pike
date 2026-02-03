resource "google_dataflow_flex_template_job" "pike" {
  provider                = google-beta
  name                    = "dataflow-flextemplates-job"
  container_spec_gcs_path = "gs://my-bucket/templates/template.json"
  parameters = {
    inputSubscription = "messages"
  }
}
