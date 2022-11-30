resource "azurerm_private_endpoint" "pike" {
  resource_group_name = "pike"
  location            = "uksouth"
  subnet_id           = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.Network/virtualNetworks/pike/subnets/pike"
  name                = "pike"
  private_service_connection {
    private_connection_resource_id = azurerm_redis_cache.pike.id
    is_manual_connection           = false
    name                           = "cachy"
    subresource_names = [
      "redisCache",
    ]
  }

  private_dns_zone_group {
    name                 = "default"
    private_dns_zone_ids = ["/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.Network/privateDnsZones/private.beer"]
  }

  timeouts {

  }
  tags = { pike = "permission" }
}
