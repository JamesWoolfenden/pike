resource "azurerm_stack_hci_extension" "pike_gen" {
  name                               = "AzureMonitorWindowsAgent"
  arc_setting_id                     = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-hci/providers/Microsoft.AzureStackHCI/clusters/hci-cl/arcSettings/default"
  publisher                          = "Microsoft.Azure.Monitor"
  type                               = "MicrosoftMonitoringAgent"
  auto_upgrade_minor_version_enabled = true
  automatic_upgrade_enabled          = true
  type_handler_version               = "1.22.0"
}
