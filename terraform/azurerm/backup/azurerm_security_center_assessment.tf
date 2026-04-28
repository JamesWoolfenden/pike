resource "azurerm_security_center_assessment" "pike_gen" {
  assessment_policy_id = "pike"
  target_resource_id   = "pike"

  status {
    code = "Healthy"
  }
}
