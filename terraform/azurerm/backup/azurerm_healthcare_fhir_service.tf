resource "azurerm_healthcare_fhir_service" "pike_gen" {
  name                = "tfexfhir"
  location            = "east us"
  resource_group_name = "tfex-resource_group"
  workspace_id        = "pike"
  kind                = "fhir-R4"

  authentication {
    authority = "https://login.microsoftonline.com/tenantId"
    audience  = "https://tfexfhir.fhir.azurehealthcareapis.com"
  }

  access_policy_object_ids = [
    data.azurerm_client_config.current.object_id
  ]

  identity {
    type = "SystemAssigned"
  }

  container_registry_login_server_url = ["tfex-container_registry_login_server"]

  cors {
    allowed_origins     = ["https://tfex.com:123", "https://tfex1.com:3389"]
    allowed_headers     = ["*"]
    allowed_methods     = ["GET", "DELETE", "PUT"]
    max_age_in_seconds  = 3600
    credentials_allowed = true
  }

  configuration_export_storage_account_name = "storage_account_name"
}
