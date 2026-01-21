
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    "compute.globalAddresses.create",
    "compute.globalAddresses.createInternal",
    "compute.globalAddresses.delete",
    "compute.globalAddresses.deleteInternal",
    "compute.globalAddresses.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "compute.networks.removePeering",
    "compute.networks.use",
    "resourcemanager.projects.get",
    "servicenetworking.services.addPeering",
    "servicenetworking.services.deleteConnection",
    "servicenetworking.services.get",
    "serviceusage.operations.get",
    "storage.buckets.get",

    //google_apigee_addons_config
    "apigee.addonsconfig.get",
    "apigee.addonsconfig.update",

    ///google_apigee_api

    //google_apigee_deployment
    "apigee.deployments.create",
    "apigee.deployments.delete",

    //google_apigee_api_product
    "apigee.apiproducts.create",
    "apigee.apiproducts.delete",
    "apigee.apiproducts.get",
    "apigee.apiproducts.update",

    //google_apigee_app_group
    "apigee.appgroups.create",
    "apigee.appgroups.delete",
    "apigee.appgroups.get",
    "apigee.appgroups.update",

    //google_apigee_developer
    "apigee.developers.create",
    "apigee.developers.delete",
    "apigee.developers.get",
    "apigee.developers.update",

    //google_apigee_developer_app
    "apigee.developerapps.create",
    "apigee.developerapps.delete",
    "apigee.developerapps.get",

    ///google_apigee_dns_zone
    "apigee.dnsZones.create",
    "apigee.dnsZones.delete",
    "apigee.dnsZones.get",

    //google_apigee_envgroup
    "apigee.envgroups.create",
    "apigee.envgroups.delete",
    "apigee.envgroups.get",
    "apigee.envgroups.update",

    //google_apigee_envgroup_attachment
    "apigee.envgroupattachments.create",
    "apigee.envgroupattachments.delete",
    "apigee.envgroupattachments.get",

    //
    "compute.globalAddresses.setLabels",

    //"google_apigee_organization"
    "apigee.organizations.create",
    "apigee.organizations.get",
    "apigee.organizations.update",
    "apigee.organizations.delete",
    "apigee.operations.get"

  ]
}
