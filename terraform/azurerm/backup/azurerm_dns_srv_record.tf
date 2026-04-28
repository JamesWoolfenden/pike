resource "azurerm_dns_srv_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300

  record {
    priority = 1
    weight   = 5
    port     = 8080
    target   = "target1.contoso.com"
  }

  tags = {
    Environment = "Production"
  }
}
