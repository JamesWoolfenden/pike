resource "google_apigee_sync_authorization" "pike" {
  name       = google_apigee_organization.org.name
  identities = ["user:james.woolfenden@gmail.com"]
}
