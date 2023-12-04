data "google_workstations_workstation_iam_policy" "pike" {
  provider               = google-beta
  workstation_cluster_id = "pike"
  workstation_config_id  = "pike"
  workstation_id         = "pike"
}
