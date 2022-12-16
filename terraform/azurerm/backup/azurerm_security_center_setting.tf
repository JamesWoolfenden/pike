resource "azurerm_security_center_setting" "pike" {
  enabled      = true
  setting_name = "MCAS"
}
