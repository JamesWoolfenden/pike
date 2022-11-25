resource "azurerm_dns_zone" "pike" {
  name                = "freebeer.site"
  resource_group_name = "pike"
  tags = {
    pike = "permissions"
  }
}
