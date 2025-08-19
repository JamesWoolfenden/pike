resource "google_dialogflow_cx_version" "version_1" {
  parent       = google_dialogflow_cx_agent.pike.start_flow
  display_name = "1.0.0"
  description  = "version 1.0.0"
}
