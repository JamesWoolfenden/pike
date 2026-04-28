resource "azurerm_container_registry_cache_rule" "pike_gen" {
  name                  = "cacherule"
  container_registry_id = "pike"
  target_repo           = "target"
  source_repo           = "docker.io/hello-world"
  credential_set_id     = "${azurerm_container_registry.acr.id}/credentialSets/example"
}
