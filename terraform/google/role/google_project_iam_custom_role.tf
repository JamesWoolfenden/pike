resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [

    "compute.networks.get",
    "compute.networks.list",
    "compute.networks.create",
    "compute.networks.delete",

    #   "dns.changes.create",
    #   "dns.changes.get",
    #   "dns.resourceRecordSets.create",
    #  "dns.resourceRecordSets.delete",
    #    "dns.resourceRecordSets.get",
    #    "dns.resourceRecordSets.list",
    #    "dns.resourceRecordSets.update",
    #    "dns.changes.list",

    #    "dns.dnsKeys.get",
    #    "dns.dnsKeys.list",

    #    "dns.managedZoneOperations.get",
    #    "dns.managedZoneOperations.list",
    #
    #    "dns.managedZones.create",
    #    "dns.managedZones.delete",
    #    "dns.managedZones.get",
    #    "dns.managedZones.getIamPolicy",
    #    "dns.managedZones.list",
    #    "dns.managedZones.update",

    "dns.networks.bindDNSResponsePolicy",
    #    "dns.networks.bindPrivateDNSPolicy",
    #    "dns.networks.bindPrivateDNSZone",
    #    "dns.networks.targetWithPeeringZone",
    #    "dns.networks.useHealthSignals",
    #
    #    "dns.policies.create",
    #    "dns.policies.delete",
    #    "dns.policies.get",
    #    "dns.policies.getIamPolicy",
    #    "dns.policies.list",
    #    "dns.policies.update",

    //"dns.projects.get",

    "dns.responsePolicies.create",
    "dns.responsePolicies.delete",
    "dns.responsePolicies.get",
    "dns.responsePolicies.list",
    "dns.responsePolicies.update",

    "dns.responsePolicyRules.create",
    "dns.responsePolicyRules.delete",
    "dns.responsePolicyRules.get",
    "dns.responsePolicyRules.list",
    "dns.responsePolicyRules.update",

    //"resourcemanager.projects.get",

  ]
}
