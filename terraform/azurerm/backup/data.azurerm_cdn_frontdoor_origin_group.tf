data "azurerm_cdn_frontdoor_origin_group" "pike_gen" {
  name                = "example-origin-group"
  profile_name        = "example-profile"
  resource_group_name = "example-resources"
}
