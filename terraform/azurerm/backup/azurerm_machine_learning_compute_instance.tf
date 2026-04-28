resource "azurerm_machine_learning_compute_instance" "pike_gen" {
  name                          = "example"
  machine_learning_workspace_id = "pike"
  virtual_machine_size          = "STANDARD_DS2_V2"
  authorization_type            = "personal"
  ssh {
    public_key = "pike"
  }
  subnet_resource_id = "pike"
  description        = "foo"
  tags = {
    foo = "bar"
  }
}
