resource "azurerm_private_dns_mx_record" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  zone_name           = "pike"
  ttl                 = 300

  record {
    preference = 10
    exchange   = "mx1.contoso.com"
  }

  record {
    preference = 20
    exchange   = "backupmx.contoso.com"
  }

  tags = {
    Environment = "Production"
  }
}
