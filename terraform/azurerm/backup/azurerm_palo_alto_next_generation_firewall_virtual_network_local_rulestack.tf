resource "azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack" "pike_gen" {
  name                = "example-ngfwvn"
  resource_group_name = "pike"
  rulestack_id        = "pike"

  network_profile {
    public_ip_address_ids = ["pike"]

    vnet_configuration {
      virtual_network_id  = "pike"
      trusted_subnet_id   = "pike"
      untrusted_subnet_id = "pike"
    }
  }
}
