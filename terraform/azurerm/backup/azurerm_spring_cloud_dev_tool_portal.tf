resource "azurerm_spring_cloud_dev_tool_portal" "pike_gen" {
  name                          = "default"
  spring_cloud_service_id       = "pike"
  public_network_access_enabled = true

  sso {
    client_id    = "example id"
    metadata_url = "https://login.microsoftonline.com/${data.azurerm_client_config.current.tenant_id}/v2.0/.well-known/openid-configuration"
    scope        = ["openid", "profile", "email"]
  }

  application_accelerator_enabled = true
  application_live_view_enabled   = true
}
