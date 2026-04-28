resource "azurerm_iothub_certificate" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  iothub_name         = "pike"
  is_verified         = true

  certificate_content = filebase64("example.cer")
}
