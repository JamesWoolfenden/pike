resource "azurerm_advisor_suppression" "pike_gen" {
  name              = "HardcodedSuppressionName"
  recommendation_id = "pike"
  resource_id       = "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  ttl               = "01:00:00:00"
}
