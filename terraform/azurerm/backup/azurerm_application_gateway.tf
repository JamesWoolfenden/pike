resource "azurerm_application_gateway" "pike_gen" {
  name                = "example-appgateway"
  resource_group_name = "pike"
  location            = "pike"

  sku {
    name     = "Standard_v2"
    tier     = "Standard_v2"
    capacity = 2
  }

  gateway_ip_configuration {
    name      = "my-gateway-ip-configuration"
    subnet_id = "pike"
  }

  frontend_port {
    name = "pike"
    port = 80
  }

  frontend_ip_configuration {
    name                 = "pike"
    public_ip_address_id = "pike"
  }

  backend_address_pool {
    name = "pike"
  }

  backend_http_settings {
    name                  = "pike"
    cookie_based_affinity = "Disabled"
    path                  = "/path1/"
    port                  = 80
    protocol              = "Http"
    request_timeout       = 60
  }

  http_listener {
    name                           = "pike"
    frontend_ip_configuration_name = "pike"
    frontend_port_name             = "pike"
    protocol                       = "Http"
  }

  request_routing_rule {
    name                       = "pike"
    priority                   = 9
    rule_type                  = "Basic"
    http_listener_name         = "pike"
    backend_address_pool_name  = "pike"
    backend_http_settings_name = "pike"
  }
}
