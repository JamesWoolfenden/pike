resource "azurerm_container_app_environment_managed_certificate" "pike_gen" {
  name                         = "example-managed-cert"
  container_app_environment_id = "pike"
  subject_name                 = "example.com"
  domain_control_validation    = "HTTP"

  depends_on = ["pike"]
}
