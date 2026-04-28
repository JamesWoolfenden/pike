resource "azurerm_spring_cloud_certificate" "pike_gen" {
  name                     = "example-scc"
  resource_group_name      = "pike"
  service_name             = "pike"
  key_vault_certificate_id = "pike"
  exclude_private_key      = true
}
