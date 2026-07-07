data "google_storage_control_project_intelligence_finding" "pike" {
  project    = "pike-477416"
  finding_id = "cross_region_egress_spike_insight_1"
}

output "google_storage_control_project_intelligence_finding" {
  value = data.google_storage_control_project_intelligence_finding.pike
}
