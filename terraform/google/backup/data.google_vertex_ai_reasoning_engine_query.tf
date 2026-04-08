data "google_vertex_ai_reasoning_engine_query" "pike" {
  provider = google-beta
}

output "google_vertex_ai_reasoning_engine_query" {
  value = data.google_vertex_ai_reasoning_engine_query.pike
}
