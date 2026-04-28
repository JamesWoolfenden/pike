resource "azurerm_arc_kubernetes_cluster" "pike_gen" {
  name                         = "example-akcc"
  resource_group_name          = "pike"
  location                     = "West Europe"
  agent_public_key_certificate = filebase64("testdata/public.cer")

  identity {
    type = "SystemAssigned"
  }

  tags = {
    ENV = "Test"
  }
}
