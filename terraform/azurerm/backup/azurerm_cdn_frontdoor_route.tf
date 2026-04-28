resource "azurerm_cdn_frontdoor_route" "pike_gen" {
  name                          = "example-route"
  cdn_frontdoor_endpoint_id     = "pike"
  cdn_frontdoor_origin_group_id = "pike"
  cdn_frontdoor_origin_ids      = ["pike"]
  cdn_frontdoor_rule_set_ids    = ["pike"]
  enabled                       = true

  forwarding_protocol    = "HttpsOnly"
  https_redirect_enabled = true
  patterns_to_match      = ["/*"]
  supported_protocols    = ["Http", "Https"]

  cdn_frontdoor_custom_domain_ids = ["pike", "pike"]
  link_to_default_domain          = false

  cache {
    query_string_caching_behavior = "IgnoreSpecifiedQueryStrings"
    query_strings                 = ["account", "settings"]
    compression_enabled           = true
    content_types_to_compress     = ["text/html", "text/javascript", "text/xml"]
  }
}
