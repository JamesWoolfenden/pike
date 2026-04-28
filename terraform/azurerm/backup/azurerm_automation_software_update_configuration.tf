resource "azurerm_automation_software_update_configuration" "pike_gen" {
  name                  = "example"
  automation_account_id = "pike"

  linux {
    classifications_included = "Security"
    excluded_packages        = ["apt"]
    included_packages        = ["vim"]
    reboot                   = "IfRequired"
  }

  pre_task {
    source = "pike"
    parameters = {
      COMPUTER_NAME = "Foo"
    }
  }

  duration = "PT2H2M2S"
}
