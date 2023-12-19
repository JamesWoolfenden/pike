data "azurerm_api_management_gateway_host_name_configuration" "pike" {
  api_management_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.ApiManagement/service/pike"
  name              = "pike"
  gateway_name      = "pike"
}
