resource "azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack" "pike_gen" {
  name                = "example-ngfwvn"
  resource_group_name = "pike"
  rulestack_id        = "pike"

  network_profile {
    public_ip_address_ids        = ["pike"]
    virtual_hub_id               = "pike"
    network_virtual_appliance_id = "pike"
  }
}
