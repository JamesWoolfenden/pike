resource "google_dialogflow_cx_environment" "pike" {
  parent       = google_dialogflow_cx_agent.pike.id
  display_name = "Development"
  description  = "Development Environment"
  version_configs {
    version = google_dialogflow_cx_version.version_1.id
  }
}
