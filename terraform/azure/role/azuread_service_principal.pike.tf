data "azuread_client_config" "current" {}

resource "azuread_application" "pike" {
  display_name = "pike"
  owners       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal" "pike" {
  application_id               = azuread_application.pike.application_id
  app_role_assignment_required = false
  owners                       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal_password" "pike" {
  service_principal_id = azuread_service_principal.pike.object_id
}

provider "azuread" {
}

provider "azurerm" {
  features {}
}

# https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/azure_cli
