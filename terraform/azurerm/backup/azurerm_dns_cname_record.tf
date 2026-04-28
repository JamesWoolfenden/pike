resource "azurerm_dns_cname_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  record              = "contoso.com"
}
