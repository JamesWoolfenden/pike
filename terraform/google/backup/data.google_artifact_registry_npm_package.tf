data "google_artifact_registry_npm_package" "pike" {
  provider = google-beta
}

output "google_artifact_registry_npm_package" {
  value = data.google_artifact_registry_npm_package.pike
}
