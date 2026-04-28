resource "azurerm_nginx_api_key" "pike_gen" {
  name                = "example-api-key"
  nginx_deployment_id = "pike"
  # We don't recommend hard coding the secret_text value. Please refer to the secret_text reference below for sources on how to manage sensitive input variables.
  secret_text   = "727c8642-6807-4254-9d02-ae93bfad21de"
  end_date_time = "2027-01-01T00:00:00Z"
}
