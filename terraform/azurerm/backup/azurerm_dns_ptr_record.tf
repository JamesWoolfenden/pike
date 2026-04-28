resource "azurerm_dns_ptr_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  records             = ["yourdomain.com"]
}
