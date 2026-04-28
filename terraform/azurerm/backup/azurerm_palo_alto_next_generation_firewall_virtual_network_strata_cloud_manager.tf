resource "azurerm_palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager" "pike_gen" {
  name                             = "example-ngfwvh"
  resource_group_name              = "pike"
  location                         = "pike"
  strata_cloud_manager_tenant_name = "example-scm-tenant"

  network_profile {
    public_ip_address_ids = ["pike"]

    vnet_configuration {
      virtual_network_id  = "pike"
      trusted_subnet_id   = "pike"
      untrusted_subnet_id = "pike"
    }
  }
}
