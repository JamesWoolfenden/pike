data "google_workstations_workstation_config_iam_policy" "pike" {
  provider               = google-beta
  workstation_cluster_id = "pike"
  workstation_config_id  = "pike"
}
