resource "azurerm_automation_watcher" "pike_gen" {
  name                           = "example"
  automation_account_id          = "pike"
  location                       = "West Europe"
  script_name                    = "pike"
  script_run_on                  = "pike"
  description                    = "example-watcher desc"
  execution_frequency_in_seconds = 42

  tags = {
    "foo" = "bar"
  }

  script_parameters = {
    foo = "bar"
  }
}
