resource "azurerm_api_connection" "pike_gen" {
  name                = "example-connection"
  resource_group_name = "pike"
  managed_api_id      = "pike"
  display_name        = "Example 1"

  parameter_values = {
    connectionString = "pike"
  }

  tags = {
    Hello = "World"
  }

  lifecycle {
    # NOTE: since the connectionString is a secure value it's not returned from the API
    ignore_changes = ["parameter_values"]
  }
}
