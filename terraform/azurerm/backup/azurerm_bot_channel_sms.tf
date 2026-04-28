resource "azurerm_bot_channel_sms" "pike_gen" {
  bot_name                        = "pike"
  location                        = "pike"
  resource_group_name             = "pike"
  sms_channel_account_security_id = "BG61f7cf5157f439b084e98256409c2815"
  phone_number                    = "+12313803556"
}
