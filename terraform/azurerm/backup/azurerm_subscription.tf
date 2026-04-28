resource "azurerm_subscription" "pike_gen" {
  subscription_name = "My Example EA Subscription"
  billing_scope_id  = "pike"
}
