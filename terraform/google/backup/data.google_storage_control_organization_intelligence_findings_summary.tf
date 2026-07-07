data "google_storage_control_organization_intelligence_findings_summary" "pike" {
  organization = "123456789012"
}

output "google_storage_control_organization_intelligence_findings_summary" {
  value = data.google_storage_control_organization_intelligence_findings_summary.pike
}
