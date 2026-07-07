data "google_storage_control_project_intelligence_findings_summary" "pike" {
  project = "pike-477416"
}

output "google_storage_control_project_intelligence_findings_summary" {
  value = data.google_storage_control_project_intelligence_findings_summary.pike
}
