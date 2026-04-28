resource "azurerm_cdn_endpoint_custom_domain" "pike_gen" {
  name            = "example-domain"
  cdn_endpoint_id = "pike"
  host_name       = "${azurerm_dns_cname_record.example.name}.${data.azurerm_dns_zone.example.name}"
}
