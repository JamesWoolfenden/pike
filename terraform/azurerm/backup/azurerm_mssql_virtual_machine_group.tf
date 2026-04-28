resource "azurerm_mssql_virtual_machine_group" "pike_gen" {
  name                = "examplegroup"
  resource_group_name = "pike"
  location            = "pike"

  sql_image_offer = "SQL2017-WS2016"
  sql_image_sku   = "Developer"

  wsfc_domain_profile {
    fqdn                = "testdomain.com"
    cluster_subnet_type = "SingleSubnet"
  }
}
