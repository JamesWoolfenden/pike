resource "azurerm_service_fabric_managed_cluster" "pike_gen" {
  name                = "example"
  resource_group_name = "example"
  location            = "West Europe"
  http_gateway_port   = 4567

  lb_rule {
    backend_port       = 38080
    frontend_port      = 80
    probe_protocol     = "http"
    probe_request_path = "/test"
    protocol           = "tcp"
  }
  client_connection_port = 12345

  node_type {
    data_disk_size_gb      = 130
    name                   = "test1"
    primary                = true
    application_port_range = "30000-49000"
    ephemeral_port_range   = "10000-20000"

    vm_size            = "Standard_DS1_v2"
    vm_image_publisher = "MicrosoftWindowsServer"
    vm_image_sku       = "2019-Datacenter-with-Containers"
    vm_image_offer     = "WindowsServer"
    vm_image_version   = "latest"
    vm_instance_count  = 5
  }
}
