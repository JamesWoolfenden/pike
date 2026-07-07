data "google_storage_control_project_intelligence_finding_revision" "pike" {
  project     = "pike-477416"
  finding_id  = "coldline_archival_storage_operations_spike_insight_1"
  revision_id = "revision_1"
}

output "google_storage_control_project_intelligence_finding_revision" {
  value = data.google_storage_control_project_intelligence_finding_revision.pike
}
