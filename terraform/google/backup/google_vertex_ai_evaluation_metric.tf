resource "google_vertex_ai_evaluation_metric" "pike" {
  evaluation_metric_id = "pike-test-metric"
  region               = "us-central1"
  display_name         = "pike-test-metric"
  metric = jsonencode({
    llmBasedMetricSpec = {
      metricPromptTemplate = "Rate the quality of the response on a scale of 1 to 5. Response: {response}"
    }
  })
}
