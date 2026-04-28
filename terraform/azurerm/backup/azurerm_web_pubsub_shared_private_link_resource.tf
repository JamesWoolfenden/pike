resource "azurerm_web_pubsub_shared_private_link_resource" "pike_gen" {
  name               = "tfex-webpubsub-splr"
  web_pubsub_id      = "pike"
  subresource_name   = "vault"
  target_resource_id = "pike"
}
