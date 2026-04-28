data "azurerm_cdn_frontdoor_profile" "pike_gen" {
  name                = "existing-cdn-profile"
  resource_group_name = "existing-resources"
}
