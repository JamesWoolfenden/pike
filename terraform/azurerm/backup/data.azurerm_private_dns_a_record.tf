data "azurerm_private_dns_a_record" "pike_gen" {
  name                = "test"
  zone_name           = "test-zone"
  resource_group_name = "test-rg"
}
