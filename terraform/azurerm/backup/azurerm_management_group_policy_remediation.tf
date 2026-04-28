resource "azurerm_management_group_policy_remediation" "pike_gen" {
  name                 = "example"
  management_group_id  = "pike"
  policy_assignment_id = "pike"
}
