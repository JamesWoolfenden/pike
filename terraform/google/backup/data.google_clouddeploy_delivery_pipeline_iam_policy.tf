data "google_clouddeploy_delivery_pipeline_iam_policy" "pike" {
  name = "pike"
}

output "google_clouddeploy_delivery_pipeline_iam_policy" {
  value = data.google_clouddeploy_delivery_pipeline_iam_policy.pike
}
