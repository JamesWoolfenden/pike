resource "azurerm_resource_deployment_script_azure_cli" "pike_gen" {
  name                = "example-rdsac"
  resource_group_name = "pike"
  location            = "West Europe"
  version             = "2.40.0"
  retention_interval  = "P1D"
  command_line        = "'foo' 'bar'"
  cleanup_preference  = "OnSuccess"
  force_update_tag    = "1"
  timeout             = "PT30M"

  script_content = <<EOF
            echo "{\"name\":{\"displayName\":\"$1 $2\"}}" > $AZ_SCRIPTS_OUTPUT_PATH
  EOF

  identity {
    type = "UserAssigned"
    identity_ids = [
      azurerm_user_assigned_identity.example.id
    ]
  }

  tags = {
    key = "value"
  }
}
