resource "azurerm_bot_channel_email" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"
  email_address       = "example.com"
}
