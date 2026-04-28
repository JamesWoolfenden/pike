resource "azurerm_email_communication_service" "pike_gen" {
  name                = "example-emailcommunicationservice"
  resource_group_name = "pike"
  data_location       = "United States"
}
