resource "azurerm_bot_channel_web_chat" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"

  site {
    name = "TestSite"
  }
}
