resource "azurerm_mssql_server" "pike_gen" {
  name                = "mssqlserver"
  resource_group_name = "pike"
  location            = "pike"
  version             = "12.0"
  administrator_login = "missadministrator"
  minimum_tls_version = "1.2"

  azuread_administrator {
    login_username = "AzureAD Admin"
    object_id      = "00000000-0000-0000-0000-000000000000"
  }

  tags = {
    environment = "production"
  }
}
