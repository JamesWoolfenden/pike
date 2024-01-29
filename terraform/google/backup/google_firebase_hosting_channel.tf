resource "google_firebase_hosting_channel" "pike" {
  provider   = google-beta
  site_id    = google_firebase_hosting_site.pike.site_id
  channel_id = "channel-basic"
}
