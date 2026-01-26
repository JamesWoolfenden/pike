resource "google_scc_event_threat_detection_custom_module" "pike" {
  organization     = "123456789"
  display_name     = "basic_custom_module"
  enablement_state = "ENABLED"
  type             = "CONFIGURABLE_BAD_IP"
  config = jsonencode({
    "metadata" : {
      "severity" : "LOW",
      "description" : "Flagged by Forcepoint as malicious",
      "recommendation" : "Contact the owner of the relevant project."
    },
    "ips" : [
      "192.0.2.1",
      "192.0.2.0/24"
    ]
  })
}
