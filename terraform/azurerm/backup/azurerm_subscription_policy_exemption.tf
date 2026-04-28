resource "azurerm_subscription_policy_exemption" "pike_gen" {
  name                 = "exampleExemption"
  subscription_id      = "pike"
  policy_assignment_id = "pike"
  exemption_category   = "Mitigated"
}
