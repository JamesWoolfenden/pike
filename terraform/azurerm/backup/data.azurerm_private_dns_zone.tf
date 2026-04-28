data "azurerm_private_dns_zone" "pike_gen" {
  name                = "contoso.internal"
  resource_group_name = "contoso-dns"
}
