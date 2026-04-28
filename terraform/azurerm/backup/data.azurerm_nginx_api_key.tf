data "azurerm_nginx_api_key" "pike_gen" {
  name                = "existing"
  nginx_deployment_id = "pike"
}
