resource "azurerm_cdn_endpoint" "pike_gen" {
  name                = "example"
  profile_name        = "pike"
  location            = "pike"
  resource_group_name = "pike"

  origin {
    name      = "example"
    host_name = "www.contoso.com"
  }
}
