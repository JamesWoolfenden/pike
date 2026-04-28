resource "azurerm_dns_aaaa_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  records             = ["2001:db8::1:0:0:1"]
}
