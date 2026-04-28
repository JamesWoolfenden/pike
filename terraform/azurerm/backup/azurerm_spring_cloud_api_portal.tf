resource "azurerm_spring_cloud_api_portal" "pike_gen" {
  name                          = "default"
  spring_cloud_service_id       = "pike"
  gateway_ids                   = ["pike"]
  https_only_enabled            = false
  public_network_access_enabled = true
  instance_count                = 1
  api_try_out_enabled           = true
  sso {
    client_id  = "test"
    issuer_uri = "https://www.example.com/issueToken"
    scope      = ["read"]
  }
}
