resource "azurerm_palo_alto_local_rulestack_rule" "pike_gen" {
  name         = "example-rule"
  rulestack_id = "pike"
  priority     = 1000
  action       = "Allow"
  protocol     = "application-default"

  applications = ["any"]

  source {
    cidrs = ["10.0.0.0/8"]
  }

  destination {
    cidrs = ["192.168.16.0/24"]
  }
}
