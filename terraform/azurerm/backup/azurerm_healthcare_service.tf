resource "azurerm_healthcare_service" "pike_gen" {
  name                = "uniquefhirname"
  resource_group_name = "sample-resource-group"
  location            = "westus2"
  kind                = "fhir-R4"
  cosmosdb_throughput = "2000"

  identity {
    type = "SystemAssigned"
  }

  access_policy_object_ids = "pike"

  configuration_export_storage_account_name = "teststorage"

  tags = {
    "environment" = "testenv"
    "purpose"     = "AcceptanceTests"
  }

  authentication_configuration {
    authority           = "https://login.microsoftonline.com/$%7Bdata.azurerm_client_config.current.tenant_id%7D"
    audience            = "https://azurehealthcareapis.com/"
    smart_proxy_enabled = "true"
  }

  cors_configuration {
    allowed_origins    = ["http://www.example.com", "http://www.example2.com"]
    allowed_headers    = ["x-tempo-*", "x-tempo2-*"]
    allowed_methods    = ["GET", "PUT"]
    max_age_in_seconds = "500"
    allow_credentials  = "true"
  }
}
