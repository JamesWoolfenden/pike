data "azurerm_cdn_frontdoor_rule_set" "pike_gen" {
  name                = "existing-rule-set"
  profile_name        = "existing-profile"
  resource_group_name = "existing-resources"
}
