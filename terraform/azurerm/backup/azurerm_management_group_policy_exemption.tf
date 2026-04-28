resource "azurerm_management_group_policy_exemption" "pike_gen" {
  name                 = "exemption1"
  management_group_id  = "pike"
  policy_assignment_id = "pike"
  exemption_category   = "Mitigated"
}
