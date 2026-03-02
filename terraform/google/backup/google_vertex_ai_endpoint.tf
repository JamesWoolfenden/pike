# Vertex AI Endpoint
resource "google_vertex_ai_endpoint" "pike" {
  name         = "pike-endpoint"
  display_name = "Pike Test Endpoint"
  location     = "europe-west2"
  project      = "pike-477416"
  description  = "Test endpoint for Pike - updated"
}
