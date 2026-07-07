data "google_storage_control_project_intelligence_findings" "pike" {
  project = "pike-477416"
}

output "google_storage_control_project_intelligence_findings" {
  value = data.google_storage_control_project_intelligence_findings.pike
}
