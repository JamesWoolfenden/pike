resource "azurerm_subscription_policy_assignment" "pike_gen" {
  name                 = "example"
  policy_definition_id = "pike"
  subscription_id      = "pike"
}
