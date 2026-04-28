resource "azurerm_private_dns_aaaa_record" "pike_gen" {
  name                = "test"
  zone_name           = "pike"
  resource_group_name = "pike"
  ttl                 = 300
  records             = ["fd5d:70bc:930e:d008:0000:0000:0000:7334", "fd5d:70bc:930e:d008::7335"]
}
