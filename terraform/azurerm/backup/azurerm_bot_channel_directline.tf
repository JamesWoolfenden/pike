resource "azurerm_bot_channel_directline" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"

  site {
    name    = "default"
    enabled = true
  }
}
