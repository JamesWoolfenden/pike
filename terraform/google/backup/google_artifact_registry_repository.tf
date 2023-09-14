resource "google_artifact_registry_repository" "pike" {
  location      = "europe-west2"
  repository_id = "pike2"
  description   = "example docker repository name"
  format        = "DOCKER"

  docker_config {
    immutable_tags = true
  }

  kms_key_name = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"

  labels = {
    "pike"    = "permissions"
    "another" = "works"
  }

}
