resource "azurerm_resource_group_policy_remediation" "pike_gen" {
  name                 = "example-policy-remediation"
  resource_group_id    = "pike"
  policy_assignment_id = "pike"
  location_filters     = ["West Europe"]
}
