resource "azurerm_vmware_express_route_authorization" "pike_gen" {
  name             = "example-authorization"
  private_cloud_id = "pike"
}
