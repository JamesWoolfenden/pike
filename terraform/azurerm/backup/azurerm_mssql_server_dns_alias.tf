resource "azurerm_mssql_server_dns_alias" "pike_gen" {
  name            = "example-dns-alias"
  mssql_server_id = "pike"
}
