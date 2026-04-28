resource "azurerm_resource_policy_remediation" "pike_gen" {
  name                 = "remediation1"
  resource_id          = "pike"
  policy_assignment_id = "pike"
}
