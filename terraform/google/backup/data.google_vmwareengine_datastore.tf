data "google_vmwareengine_datastore" "pike" {
  location = "us-central1"
  name     = "pike"
}

output "google_vmwareengine_datastore" {
  value = data.google_vmwareengine_datastore.pike
}
