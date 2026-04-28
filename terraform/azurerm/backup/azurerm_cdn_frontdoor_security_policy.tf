resource "azurerm_cdn_frontdoor_security_policy" "pike_gen" {
  name                     = "Example-Security-Policy"
  cdn_frontdoor_profile_id = "pike"

  security_policies {
    firewall {
      cdn_frontdoor_firewall_policy_id = "pike"

      association {
        domain {
          cdn_frontdoor_domain_id = "pike"
        }
        patterns_to_match = ["/*"]
      }
    }
  }
}
