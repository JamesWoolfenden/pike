resource "azurerm_bot_channel_alexa" "pike_gen" {
  bot_name            = "pike"
  location            = "pike"
  resource_group_name = "pike"
  skill_id            = "amzn1.ask.skill.00000000-0000-0000-0000-000000000000"
}
