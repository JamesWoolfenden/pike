data "azurerm_cdn_frontdoor_firewall_policy" "pike_gen" {
  name                = "examplecdnfdwafpolicy"
  resource_group_name = "pike"
}
