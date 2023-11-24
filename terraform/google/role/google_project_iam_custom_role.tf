resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //ggoogle_compute_regions
    "compute.regions.list",
    //google_compute_resource_policy
    "compute.resourcePolicies.get",
    //google_compute_router, google_compute_router_nat,google_compute_router_status
    "compute.routers.get",
    //google_compute_snapshot
    "compute.snapshots.get",
    //google_compute_snapshot_iam_policy
    "compute.snapshots.getIamPolicy",
    //google_compute_ssl_certificate
    "compute.sslCertificates.get",
    //google_compute_ssl_policy
    "compute.sslPolicies.get",
    //google_compute_subnetwork_iam_policy
    "compute.subnetworks.getIamPolicy",
    //google_compute_vpn_gateway
    "compute.targetVpnGateways.get"

  ]
}
