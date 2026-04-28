resource "azurerm_route_filter" "pike_gen" {
  name                = "example"
  resource_group_name = "example"
  location            = "East US"

  rule {
    name        = "rule"
    access      = "Allow"
    rule_type   = "Community"
    communities = ["12076:52004"]
  }
}
