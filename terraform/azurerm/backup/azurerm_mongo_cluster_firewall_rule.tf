resource "azurerm_mongo_cluster_firewall_rule" "pike_gen" {
  name             = "example-firewall-rule"
  mongo_cluster_id = "pike"
  start_ip_address = "10.0.0.1"
  end_ip_address   = "10.0.0.255"
}
