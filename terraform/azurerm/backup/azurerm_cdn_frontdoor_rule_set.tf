resource "azurerm_cdn_frontdoor_rule_set" "pike_gen" {
  name                     = "ExampleRuleSet"
  cdn_frontdoor_profile_id = "pike"
}
