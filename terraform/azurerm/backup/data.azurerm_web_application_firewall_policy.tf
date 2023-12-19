data "azurerm_web_application_firewall_policy" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
