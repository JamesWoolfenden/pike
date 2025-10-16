resource "google_vertex_ai_endpoint_with_model_garden_deployment" "deploy" {
  publisher_model_name = "publishers/google/models/paligemma@paligemma-224-float32"
  location             = "europe-west2"
  model_config {
    accept_eula = true
  }
}
