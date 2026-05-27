
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    //google_artifact_registry_file
    "artifactregistry.files.get",

    //google_compute_region_instant_snapshot_iam_policy
    "compute.instantSnapshots.getIamPolicy",

    //google_logging_log_view
    # non required

    //google_network_connectivity_hub_iam_policy
    "networkconnectivity.hubs.getIamPolicy",

    //google_data_lineage_config
    "datalineage.configs.get"



  ]
}
