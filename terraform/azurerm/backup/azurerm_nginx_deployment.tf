resource "azurerm_nginx_deployment" "pike_gen" {
  name                      = "example-nginx"
  resource_group_name       = "pike"
  sku                       = "standardv3_Monthly"
  location                  = "pike"
  automatic_upgrade_channel = "stable"

  frontend_public {
    ip_address = ["pike"]
  }
  network_interface {
    subnet_id = "pike"
  }

  capacity = 20

  email = "user@test.com"
}
