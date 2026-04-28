data "azurerm_cdn_frontdoor_custom_domain" "pike_gen" {
  name                = "pike"
  profile_name        = "pike"
  resource_group_name = "pike"
}
