resource "azurerm_kusto_cluster_managed_private_endpoint" "pike_gen" {
  name                         = "examplempe"
  resource_group_name          = "pike"
  cluster_name                 = "pike"
  private_link_resource_id     = "pike"
  private_link_resource_region = "pike"
  group_id                     = "blob"
  request_message              = "Please Approve"
}
