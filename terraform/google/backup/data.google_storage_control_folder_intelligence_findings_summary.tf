data "google_storage_control_folder_intelligence_findings_summary" "pike" {
  folder = "123456789"
}

output "google_storage_control_folder_intelligence_findings_summary" {
  value = data.google_storage_control_folder_intelligence_findings_summary.pike
}
