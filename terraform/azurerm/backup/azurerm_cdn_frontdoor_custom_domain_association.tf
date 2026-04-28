resource "azurerm_cdn_frontdoor_custom_domain_association" "pike_gen" {
  cdn_frontdoor_custom_domain_id = "pike"
  cdn_frontdoor_route_ids        = ["pike"]
}
