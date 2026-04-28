resource "azurerm_virtual_machine_extension" "pike_gen" {
  name                 = "hostname"
  virtual_machine_id   = "pike"
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.0"

  settings = <<SETTINGS
 {
  "commandToExecute": "hostname && uptime"
 }
SETTINGS


  tags = {
    environment = "Production"
  }
}
