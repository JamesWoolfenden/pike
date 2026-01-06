
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    "certificatemanager.dnsauthorizations.create",
    "certificatemanager.dnsauthorizations.delete",
    "certificatemanager.dnsauthorizations.get",
    "certificatemanager.dnsauthorizations.update",
    "certificatemanager.operations.delete",
    "privateca.caPools.create",
    "privateca.caPools.delete",
    "privateca.caPools.get",
    "privateca.caPools.update",
    "privateca.certificateAuthorities.create",
    "privateca.certificateAuthorities.delete",
    "privateca.certificateAuthorities.get",
    "privateca.certificateAuthorities.update",
    "storage.buckets.get",

    //google_privateca_ca_pool
    "privateca.operations.get",


    //google_certificate_manager_certificate_map
    "certificatemanager.certmaps.create",


    //google_certificate_manager_dns_authorization
    "certificatemanager.operations.get",


    //google_active_directory_domain
    "managedidentities.domains.create",
    "managedidentities.domains.get",
    "managedidentities.domains.update",
    "managedidentities.domains.delete",
    "managedidentities.operations.get",

    //google_active_directory_domain_trust
    "managedidentities.domains.attachTrust",


    //google_active_directory_peering
    "managedidentities.peerings.create",

    //google_certificate_manager_certificate_map
    "certificatemanager.certmaps.get",
    "certificatemanager.certmaps.delete",


    //google_certificate_manager_certificate_issuance_config
    "certificatemanager.certissuanceconfigs.create",


    //google_certificate_manager_certificate
    "certificatemanager.certs.create",
    "certificatemanager.dnsauthorizations.use",
    "certificatemanager.certs.get",
    "certificatemanager.certs.delete",

    //google_certificate_manager_certificate_map_entry
    "certificatemanager.certmapentries.create",
    "certificatemanager.certmapentries.delete",
    "certificatemanager.certmapentries.get",
    "certificatemanager.certmapentries.update",
    "certificatemanager.certs.use"

  ]
}
