data "google_vmwareengine_upgrades" "pike" {
  provider = google-beta
  parent   = "projects/pike-477416/locations/us-west1-a"
}

output "google_vmwareengine_upgrades" {
  value = data.google_vmwareengine_upgrades.pike
}
