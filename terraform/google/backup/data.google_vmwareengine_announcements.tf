data "google_vmwareengine_announcements" "pike" {
  provider = google-beta
  parent   = "projects/pike-477416/locations/us-central1-a"
}

output "google_vmwareengine_announcements" {
  value = data.google_vmwareengine_announcements.pike
}
