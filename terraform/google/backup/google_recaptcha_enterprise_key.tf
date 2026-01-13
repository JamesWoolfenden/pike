resource "google_recaptcha_enterprise_key" "pike" {
  display_name = "pike"
  # project      = "my-project-name"

  web_settings {
    integration_type  = "SCORE"
    allow_all_domains = true
  }

  labels = {}
}
