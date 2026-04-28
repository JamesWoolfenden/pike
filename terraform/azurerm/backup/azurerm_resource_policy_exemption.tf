resource "azurerm_resource_policy_exemption" "pike_gen" {
  name                 = "exemption1"
  resource_id          = "pike"
  policy_assignment_id = "pike"
  exemption_category   = "Mitigated"
}
