resource "azurerm_palo_alto_next_generation_firewall_virtual_network_panorama" "pike_gen" {
  name                   = "example-ngfwvh"
  resource_group_name    = "pike"
  location               = "pike"
  panorama_base64_config = "e2RnbmFtZTogY25nZnctYXotZXhhbXBsZSwgdHBsbmFtZTogY25nZnctZXhhbXBsZS10ZW1wbGF0ZS1zdGFjaywgZXhhbXBsZS1wYW5vcmFtYS1zZXJ2ZXI6IDE5Mi4xNjguMC4xLCB2bS1hdXRoLWtleTogMDAwMDAwMDAwMDAwMDAwLCBleHBpcnk6IDIwMjQvMDcvMzF9Cg=="

  network_profile {
    public_ip_address_ids = ["pike"]

    vnet_configuration {
      virtual_network_id  = "pike"
      trusted_subnet_id   = "pike"
      untrusted_subnet_id = "pike"
    }
  }
}
