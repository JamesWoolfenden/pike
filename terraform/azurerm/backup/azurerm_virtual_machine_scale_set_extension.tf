resource "azurerm_virtual_machine_scale_set_extension" "pike_gen" {
  name                         = "example"
  virtual_machine_scale_set_id = "pike"
  publisher                    = "Microsoft.Azure.Extensions"
  type                         = "CustomScript"
  type_handler_version         = "2.0"
  settings = jsonencode({
    "commandToExecute" = "echo $HOSTNAME"
  })
}
