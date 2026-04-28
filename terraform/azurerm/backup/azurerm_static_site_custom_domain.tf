resource "azurerm_static_site_custom_domain" "pike_gen" {
  static_site_id  = "pike"
  domain_name     = "${azurerm_dns_cname_record.example.name}.${azurerm_dns_cname_record.example.zone_name}"
  validation_type = "cname-delegation"
}
