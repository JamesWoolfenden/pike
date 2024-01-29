resource "google_firebase_hosting_release" "pike" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.pike.site_id
  version_name = google_firebase_hosting_version.pike.name
}
