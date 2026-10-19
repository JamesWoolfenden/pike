data "google_eventarc_pipeline_iam_policy" "pike" {
  pipeline_id = "pike-test-pipeline"
}

output "google_eventarc_pipeline_iam_policy" {
  value = data.google_eventarc_pipeline_iam_policy.pike
}