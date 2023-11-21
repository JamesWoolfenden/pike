resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_compute_global_address
    "compute.globalAddresses.get",
    //google_compute_forwarding_rule
    "compute.forwardingRules.get",
    //google_compute_default_service_account
    "compute.projects.get",
    //google_compute_backend_service_iam_policy
    "compute.backendServices.getIamPolicy",
    //google_compute_backend_service
    "compute.backendServices.get",
    //google_compute_backend_bucket_iam_policy
    "compute.backendBuckets.getIamPolicy",
    //google_compute_backend_bucket
    "compute.backendBuckets.get",
    //google_compute_addresses
    "compute.addresses.list",
    //google_compute_address
    "compute.addresses.get"
  ]
}
