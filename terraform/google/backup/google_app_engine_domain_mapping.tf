resource "google_app_engine_domain_mapping" "pike" {
  domain_name = "verified-domain.com"

  ssl_settings {
    ssl_management_type = "AUTOMATIC"
  }

}
