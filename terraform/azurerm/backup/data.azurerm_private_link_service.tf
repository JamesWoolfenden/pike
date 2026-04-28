data "azurerm_private_link_service" "pike_gen" {
  name                = "myPrivateLinkService"
  resource_group_name = "PrivateLinkServiceRG"
}
