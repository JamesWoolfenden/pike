data "azurerm_monitor_scheduled_query_rules_log" "pike_gen" {
  resource_group_name = "terraform-example-rg"
  name                = "tfex-queryrule"
}
