resource "google_apigee_environment_api_revision_deployment" "pike" {
  org_id            = "my-org"
  environment       = "dev"
  api               = "hello-proxy"
  revision          = 1
  override          = true
  sequenced_rollout = true
}
