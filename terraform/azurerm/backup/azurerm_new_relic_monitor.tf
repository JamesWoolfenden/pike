resource "azurerm_new_relic_monitor" "pike_gen" {
  name                = "example-nrm"
  resource_group_name = "pike"
  location            = "pike"

  plan {
    effective_date = "2023-06-06T00:00:00Z"
  }

  user {
    email        = "user@example.com"
    first_name   = "Example"
    last_name    = "User"
    phone_number = "+12313803556"
  }

  identity {
    type = "SystemAssigned"
  }
}
