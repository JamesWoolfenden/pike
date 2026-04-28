data "azurerm_marketplace_agreement" "pike_gen" {
  publisher = "barracudanetworks"
  offer     = "waf"
  plan      = "hourly"
}
