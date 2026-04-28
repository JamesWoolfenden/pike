data "azurerm_dns_mx_record" "pike_gen" {
  name                = "test"
  zone_name           = "test-zone"
  resource_group_name = "test-rg"
}
