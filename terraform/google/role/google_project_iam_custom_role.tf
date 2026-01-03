
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    //google_compute_region_security_policy
    "compute.regionSecurityPolicies.get",

    //google_compute_storage_pool
    "compute.storagePools.get",

    //google_gke_hub_membership_binding
    "gkehub.membershipbindings.get",

    //google_service_networking_peered_dns_domain
    "servicenetworking.services.listPeeredDnsDomains"
  ]
}
