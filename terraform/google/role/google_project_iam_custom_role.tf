
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    //google_scc_event_threat_detection_custom_module
    "securitycenter.eventthreatdetectioncustommodules.create",
    "securitycenter.eventthreatdetectioncustommodules.delete",
    "securitycenter.eventthreatdetectioncustommodules.get",
    "securitycenter.eventthreatdetectioncustommodules.update",

    //google_scc_project_custom_module
    "securitycenter.securityhealthanalyticscustommodules.create",
    "securitycenter.securityhealthanalyticscustommodules.delete",
    "securitycenter.securityhealthanalyticscustommodules.get",
    "securitycenter.securityhealthanalyticscustommodules.update"

  ]
}
