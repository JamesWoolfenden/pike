data "google_iam_workload_identity_pool_openid_config" "pike" {
  resource_name = "https://iam.googleapis.com/projects/pike-477416/locations/global/workloadIdentityPools/pike-test-pool/providers/pike-test-provider"
}

output "google_iam_workload_identity_pool_openid_config" {
  value = data.google_iam_workload_identity_pool_openid_config.pike
}