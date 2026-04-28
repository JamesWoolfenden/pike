data "azurerm_cdn_frontdoor_endpoint" "pike_gen" {
  name                = "existing-endpoint"
  profile_name        = "existing-cdn-profile"
  resource_group_name = "existing-resources"
}
