resource "azurerm_vpn_server_configuration_policy_group" "pike_gen" {
  name                        = "example-VPNSCPG"
  vpn_server_configuration_id = "pike"

  policy {
    name  = "policy1"
    type  = "RadiusAzureGroupId"
    value = "6ad1bd08"
  }
}
