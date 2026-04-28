resource "azurerm_private_dns_ptr_record" "pike_gen" {
  name                = "15"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  records             = ["test.example.com"]
}
