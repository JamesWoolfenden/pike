resource "azurerm_ip_group_cidr" "pike_gen" {
  ip_group_id = "pike"
  cidr        = "10.10.10.0/24"
}
