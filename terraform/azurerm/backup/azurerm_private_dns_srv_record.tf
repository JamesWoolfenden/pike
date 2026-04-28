resource "azurerm_private_dns_srv_record" "pike_gen" {
  name                = "test"
  resource_group_name = "pike"
  zone_name           = "pike"
  ttl                 = 300

  record {
    priority = 1
    weight   = 5
    port     = 8080
    target   = "target1.contoso.com"
  }

  record {
    priority = 10
    weight   = 10
    port     = 8080
    target   = "target2.contoso.com"
  }

  tags = {
    Environment = "Production"
  }
}
