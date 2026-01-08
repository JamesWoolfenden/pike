resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "example-pool"
  mode                      = "TRUST_DOMAIN"
}

# resource "google_iam_workload_identity_pool_namespace" "ns" {
#   provider = google-beta
#
#   workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
#   workload_identity_pool_namespace_id = "example-namespace"
# }

resource "google_iam_workload_identity_pool_managed_identity" "example" {
  provider = google-beta

  workload_identity_pool_id                  = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id        = "23423423423"
  workload_identity_pool_managed_identity_id = "example-managed-identity"
}
