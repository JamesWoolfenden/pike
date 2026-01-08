
resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "example-namespace"
}
