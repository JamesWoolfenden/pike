resource "azurerm_container_registry_credential_set" "pike_gen" {
  name                  = "exampleCredentialSet"
  container_registry_id = "pike"
  login_server          = "docker.io"
  identity {
    type = "SystemAssigned"
  }
  authentication_credentials {
    username_secret_id = "https://example-keyvault.vault.azure.net/secrets/example-user-name"
    password_secret_id = "https://example-keyvault.vault.azure.net/secrets/example-user-password"
  }
}
