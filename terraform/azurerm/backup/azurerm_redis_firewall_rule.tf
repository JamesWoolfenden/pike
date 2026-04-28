resource "azurerm_redis_firewall_rule" "pike_gen" {
  name                = "someIPrange"
  redis_cache_name    = "pike"
  resource_group_name = "pike"
  start_ip            = "1.2.3.4"
  end_ip              = "2.3.4.5"
}
