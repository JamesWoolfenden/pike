resource "google_firebase_hosting_version" "pike" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.pike.site_id
  config {
    redirects {
      glob        = "/google/**"
      status_code = 302
      location    = "https://www.google.com"
    }
  }
}
