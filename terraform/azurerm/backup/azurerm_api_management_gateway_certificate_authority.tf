resource "azurerm_api_management_gateway_certificate_authority" "pike_gen" {
  api_management_id = "pike"
  certificate_name  = "pike"
  gateway_name      = "pike"
  is_trusted        = true
}
