resource "azurerm_signalr_shared_private_link_resource" "pike_gen" {
  name               = "tfex-signalr-splr"
  signalr_service_id = "pike"
  sub_resource_name  = "vault"
  target_resource_id = "pike"
}
