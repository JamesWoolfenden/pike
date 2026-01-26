resource "google_app_engine_application_url_dispatch_rules" "pike" {
  dispatch_rules {
    domain  = "*"
    path    = "/*"
    service = "default"
  }

  dispatch_rules {
    domain  = "*"
    path    = "/admin/*"
    service = "fred"
  }
}
