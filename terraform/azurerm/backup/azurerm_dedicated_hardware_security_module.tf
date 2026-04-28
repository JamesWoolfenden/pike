resource "azurerm_dedicated_hardware_security_module" "pike_gen" {
  name                = "example-hsm"
  location            = "pike"
  resource_group_name = "pike"
  sku_name            = "payShield10K_LMK1_CPS60"

  management_network_profile {
    network_interface_private_ip_addresses = ["10.2.1.7"]
    subnet_id                              = "pike"
  }

  network_profile {
    network_interface_private_ip_addresses = ["10.2.1.8"]
    subnet_id                              = "pike"
  }

  stamp_id = "stamp2"

  tags = {
    env = "Test"
  }

  depends_on = ["pike"]
}
