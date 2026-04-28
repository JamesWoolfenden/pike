resource "azurerm_virtual_machine_run_command" "pike_gen" {
  name               = "example-vmrc"
  location           = "pike"
  virtual_machine_id = "pike"
  source {
    script = "echo 'hello world'"
  }
}
