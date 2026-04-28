resource "azurerm_search_shared_private_link_service" "pike_gen" {
  name               = "example-spl"
  search_service_id  = "pike"
  subresource_name   = "blob"
  target_resource_id = "pike"
  request_message    = "please approve"
}
