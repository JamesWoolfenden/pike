resource "azurerm_cdn_frontdoor_endpoint" "pike_gen" {
  name                     = "example-endpoint"
  cdn_frontdoor_profile_id = "pike"

  tags = {
    ENV = "example"
  }
}
