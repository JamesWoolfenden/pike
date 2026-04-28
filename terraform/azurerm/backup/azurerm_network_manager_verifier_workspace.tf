resource "azurerm_network_manager_verifier_workspace" "pike_gen" {
  name               = "example"
  network_manager_id = "pike"
  location           = "pike"
  description        = "This is an example verifier workspace"

  tags = {
    foo = "bar"
    env = "example"
  }
}
