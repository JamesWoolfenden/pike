resource "azurerm_bot_channel_slack" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"
  client_id           = "exampleId"
}
