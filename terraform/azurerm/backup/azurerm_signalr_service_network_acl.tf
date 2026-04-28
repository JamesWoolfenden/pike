resource "azurerm_signalr_service_network_acl" "pike_gen" {
  signalr_service_id = "pike"
  default_action     = "Deny"

  public_network {
    allowed_request_types = ["ClientConnection"]
  }

  private_endpoint {
    id                    = "pike"
    allowed_request_types = ["ServerConnection"]
  }
}
