resource "azurerm_analysis_services_server" "pike_gen" {
  name                     = "analysisservicesserver"
  location                 = "pike"
  resource_group_name      = "pike"
  sku                      = "S0"
  admin_users              = ["myuser@domain.tld"]
  power_bi_service_enabled = true

  ipv4_firewall_rule {
    name        = "myRule1"
    range_start = "210.117.252.0"
    range_end   = "210.117.252.255"
  }

  tags = {
    abc = 123
  }
}
