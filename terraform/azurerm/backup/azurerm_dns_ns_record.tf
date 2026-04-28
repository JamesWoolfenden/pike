resource "azurerm_dns_ns_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300

  records = ["ns1.contoso.com.", "ns2.contoso.com."]

  tags = {
    Environment = "Production"
  }
}
