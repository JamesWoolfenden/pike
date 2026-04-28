data "azurerm_nginx_certificate" "pike_gen" {
  name                = "existing"
  nginx_deployment_id = "pike"
}
