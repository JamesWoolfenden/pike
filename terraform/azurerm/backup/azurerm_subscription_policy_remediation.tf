resource "azurerm_subscription_policy_remediation" "pike_gen" {
  name                 = "example"
  subscription_id      = "pike"
  policy_assignment_id = "pike"
}
