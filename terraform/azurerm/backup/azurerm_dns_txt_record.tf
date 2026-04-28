resource "azurerm_dns_txt_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300

  record {
    value = "google-site-authenticator"
  }

  record {
    value = "more site information here"
  }

  tags = {
    Environment = "Production"
  }
}
