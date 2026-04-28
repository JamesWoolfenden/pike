resource "azurerm_stream_analytics_managed_private_endpoint" "pike_gen" {
  name                          = "exampleprivateendpoint"
  resource_group_name           = "pike"
  stream_analytics_cluster_name = "pike"
  target_resource_id            = "pike"
  subresource_name              = "blob"
}
