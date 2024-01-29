resource "google_firebase_hosting_custom_domain" "pike" {
  provider      = google-beta
  site_id       = google_firebase_hosting_site.pike.site_id
  custom_domain = "custom.domain.com"
}
