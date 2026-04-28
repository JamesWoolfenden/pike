resource "azurerm_cdn_frontdoor_origin" "pike_gen" {
  name                          = "example-origin"
  cdn_frontdoor_origin_group_id = "pike"
  enabled                       = true

  certificate_name_check_enabled = false

  host_name          = "contoso.com"
  http_port          = 80
  https_port         = 443
  origin_host_header = "www.contoso.com"
  priority           = 1
  weight             = 1
}
