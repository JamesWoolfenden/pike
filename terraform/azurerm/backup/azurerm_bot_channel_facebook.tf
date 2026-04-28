resource "azurerm_bot_channel_facebook" "pike_gen" {
  bot_name                = "pike"
  location                = "pike"
  resource_group_name     = "pike"
  facebook_application_id = "563490254873576"

  page {
    id = "876248795081953"
  }
}
