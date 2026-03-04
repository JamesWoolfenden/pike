resource "google_cloud_quotas_quota_adjuster_settings" "adjuster_settings" {
  provider   = google-beta
  parent     = "projects/104740170505"
  enablement = "ENABLED"
}
