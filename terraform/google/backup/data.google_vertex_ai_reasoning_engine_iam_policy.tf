data "google_vertex_ai_reasoning_engine_iam_policy" "pike" {
  provider = google-beta
}

output "google_vertex_ai_reasoning_engine_iam_policy" {
  value = data.google_vertex_ai_reasoning_engine_iam_policy.pike
}
