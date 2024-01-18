resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_workbench_instance_iam_policy
    "notebooks.instances.getIamPolicy",
    //google_vmwareengine_vcenter_credentials
    "vmwareengine.privateClouds.showVcenterCredentials",

    //google_vmwareengine_external_address
    "vmwareengine.externalAddresses.get",
    //google_vmwareengine_nsx_credentials
    "vmwareengine.privateClouds.showNsxCredentials",
    //google_vmwareengine_subnet
    "vmwareengine.subnets.get",

    //google_compute_region_disk
    "compute.disks.get",

    //google_compute_reservation
    "compute.reservations.get",

    //google_filestore_instance
    "file.instances.get",

    //google_logging_project_settings, google_logging_project_settings
    "logging.settings.get"
  ]
}
