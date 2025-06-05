resource "google_iam_workload_identity_pool_provider" "pike" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pike.workload_identity_pool_id
  workload_identity_pool_provider_id = "example-prvdr"
  attribute_mapping = {
    "google.subject" = "assertion.sub"
  }
  oidc {
    issuer_uri = "https://sts.windows.net/azure-tenant-id"
  }
  display_name = "Example IDP provider"
}
