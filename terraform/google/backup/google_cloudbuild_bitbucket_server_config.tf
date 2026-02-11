resource "google_cloudbuild_bitbucket_server_config" "bbs-config" {
  config_id = "bbs-config"
  location  = "us-central1"
  host_uri  = "https://bbs.com"
  secrets {
    admin_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
    read_access_token_version_name  = "projects/myProject/secrets/mybbspat/versions/1"
    webhook_secret_version_name     = "projects/myProject/secrets/mybbspat/versions/1"
  }
  username = "test"
  api_key  = "<api-key>"
}
