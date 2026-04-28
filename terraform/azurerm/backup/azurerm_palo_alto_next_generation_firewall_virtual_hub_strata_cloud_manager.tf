resource "azurerm_palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager" "pike_gen" {
  name                             = "example"
  resource_group_name              = "example"
  location                         = "West Europe"
  strata_cloud_manager_tenant_name = "example"

  network_profile {
    public_ip_address_ids        = ["pike"]
    virtual_hub_id               = "pike"
    network_virtual_appliance_id = "pike"
  }
}
