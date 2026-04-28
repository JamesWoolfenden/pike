resource "azurerm_dns_a_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  records             = ["10.0.180.17"]
}
