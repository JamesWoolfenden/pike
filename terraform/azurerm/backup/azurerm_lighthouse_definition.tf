resource "azurerm_lighthouse_definition" "pike_gen" {
  name               = "Sample definition"
  description        = "This is a lighthouse definition created via Terraform"
  managing_tenant_id = "00000000-0000-0000-0000-000000000000"
  scope              = "/subscriptions/00000000-0000-0000-0000-000000000000"

  authorization {
    principal_id           = "00000000-0000-0000-0000-000000000000"
    role_definition_id     = "pike"
    principal_display_name = "Tier 1 Support"
  }
}
