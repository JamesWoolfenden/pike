resource "azurerm_nginx_certificate" "pike_gen" {
  name                     = "examplecert"
  nginx_deployment_id      = "pike"
  key_virtual_path         = "/src/cert/soservermekey.key"
  certificate_virtual_path = "/src/cert/server.cert"
  key_vault_secret_id      = "pike"
}
