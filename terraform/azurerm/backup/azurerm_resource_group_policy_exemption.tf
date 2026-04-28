resource "azurerm_resource_group_policy_exemption" "pike_gen" {
  name                 = "exampleExemption"
  resource_group_id    = "pike"
  policy_assignment_id = "pike"
  exemption_category   = "Mitigated"
}
