
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    // google_workstations_workstation

    // google_workstations_workstation_cluster


    // google_workstations_workstation_config


    // google_workflows_workflow


    //google_vmwareengine_cluster
    "vmwareengine.clusters.create",
    "vmwareengine.clusters.delete",
    "vmwareengine.clusters.get",
    "vmwareengine.clusters.update",

    //google_vmwareengine_external_access_rule
    "vmwareengine.externalAccessRules.create",
    "vmwareengine.externalAccessRules.delete",
    "vmwareengine.externalAccessRules.get",
    "vmwareengine.externalAccessRules.update",

    //google_vmwareengine_external_address

    //google_vmwareengine_network

    //google_vmwareengine_network_peering

    //google_vmwareengine_network_policy

    //google_vmwareengine_private_cloud



    //google_vmwareengine_subnet

  ]
}
