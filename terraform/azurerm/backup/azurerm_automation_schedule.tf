resource "azurerm_automation_schedule" "pike_gen" {
  name                    = "tfex-automation-schedule"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  frequency               = "Week"
  interval                = 1
  timezone                = "Australia/Perth"
  start_time              = "2014-04-15T18:00:15+02:00"
  description             = "This is an example schedule"
  week_days               = ["Friday"]
}
