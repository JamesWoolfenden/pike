resource "azurerm_web_pubsub_network_acl" "pike_gen" {
  web_pubsub_id  = "pike"
  default_action = "Allow"
  public_network {
    denied_request_types = ["ClientConnection"]
  }

  private_endpoint {
    id                   = "pike"
    denied_request_types = ["RESTAPI", "ClientConnection"]
  }

  depends_on = [
    azurerm_private_endpoint.example
  ]
}
