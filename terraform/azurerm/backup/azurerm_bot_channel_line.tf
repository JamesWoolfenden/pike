resource "azurerm_bot_channel_line" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"

  line_channel {
    secret = "aagfdgfd123567"
  }
}
