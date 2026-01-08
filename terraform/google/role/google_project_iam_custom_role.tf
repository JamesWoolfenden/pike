
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "compute.addresses.create",
    "compute.addresses.delete",
    "compute.addresses.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "compute.networks.updatePolicy",
    "compute.networks.use",
    "compute.subnetworks.create",
    "compute.subnetworks.delete",
    "compute.subnetworks.get",
    "compute.subnetworks.update",
    "compute.subnetworks.use",
    "storage.buckets.get",


    //google_compute_route
    "compute.routes.create",
    "compute.routes.delete",
    "compute.routes.get",

    //google_compute_router_interface google_compute_router_peer
    "compute.routers.create",
    "compute.routers.delete",
    "compute.routers.get",
    "compute.routers.update",

    //google_compute_network
    "compute.globalOperations.get",

    //google_compute_address
    "compute.addresses.setLabels",

    //google_compute_router_route_policy
    "compute.routers.updateRoutePolicy",
    "compute.routers.getRoutePolicy",
    "compute.routers.deleteRoutePolicy",

    //google_compute_router_interface
    "compute.vpnTunnels.get"
  ]
}
