data "azurerm_cdn_frontdoor_secret" "pike_gen" {
  name                = "example-secret"
  profile_name        = "example-profile"
  resource_group_name = "example-resources"
}
