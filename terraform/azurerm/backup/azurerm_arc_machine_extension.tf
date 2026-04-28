resource "azurerm_arc_machine_extension" "pike_gen" {
  name           = "example"
  location       = "West Europe"
  arc_machine_id = "pike"
  publisher      = "Microsoft.Azure.Monitor"
  type           = "AzureMonitorLinuxAgent"
}
