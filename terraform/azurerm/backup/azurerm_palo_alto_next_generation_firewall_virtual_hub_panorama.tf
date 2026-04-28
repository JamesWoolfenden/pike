resource "azurerm_palo_alto_next_generation_firewall_virtual_hub_panorama" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"

  network_profile {
    public_ip_address_ids        = ["pike"]
    virtual_hub_id               = "pike"
    network_virtual_appliance_id = "pike"
  }

  panorama_base64_config = "VGhpcyBpcyBub3QgYSByZWFsIGNvbmZpZywgcGxlYXNlIHVzZSB5b3VyIFBhbm9yYW1hIHNlcnZlciB0byBnZW5lcmF0ZSBhIHJlYWwgdmFsdWUgZm9yIHRoaXMgcHJvcGVydHkhCg=="
}
