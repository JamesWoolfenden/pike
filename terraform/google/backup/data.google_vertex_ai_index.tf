data "google_vertex_ai_index" "pike" {
  name   = "pike"
  region = "us-central1"
}

output "google_vertex_ai_index" {
  value = data.google_vertex_ai_index.pike
}
