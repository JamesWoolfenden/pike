resource "google_apigee_control_plane_access" "pike" {
  name = "pike"
  synchronizer_identities = [
    "user:james.woolfenden@gmail.com",
  ]
  analytics_publisher_identities = [
    "user:james.woolfenden@gmail.com",
  ]
}
