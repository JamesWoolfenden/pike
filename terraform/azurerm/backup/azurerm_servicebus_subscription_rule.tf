resource "azurerm_servicebus_subscription_rule" "pike_gen" {
  name            = "tfex_servicebus_rule"
  subscription_id = "pike"
  filter_type     = "SqlFilter"
  sql_filter      = "colour = 'red'"
}
