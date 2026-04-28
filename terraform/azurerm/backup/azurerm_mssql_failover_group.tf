resource "azurerm_mssql_failover_group" "pike_gen" {
  name      = "example"
  server_id = "pike"
  databases = [
    azurerm_mssql_database.example.id
  ]

  partner_server {
    id = "pike"
  }

  read_write_endpoint_failover_policy {
    mode          = "Automatic"
    grace_minutes = 80
  }

  tags = {
    environment = "prod"
    database    = "example"
  }
}
