resource "azurerm_private_dns_txt_record" "pike_gen" {
  name                = "test"
  resource_group_name = "pike"
  zone_name           = "pike"
  ttl                 = 300

  record {
    value = "v=spf1 mx ~all"
  }
}
