data "azurerm_management_group_template_deployment" "pike_gen" {
  name                = "existing"
  management_group_id = "00000000-0000-0000-000000000000"
}
