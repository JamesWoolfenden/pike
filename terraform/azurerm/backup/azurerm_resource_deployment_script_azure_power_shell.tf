resource "azurerm_resource_deployment_script_azure_power_shell" "pike_gen" {
  name                = "example-rdsaps"
  resource_group_name = "pike"
  location            = "West Europe"
  version             = "8.3"
  retention_interval  = "P1D"
  command_line        = "-name \"John Dole\""
  cleanup_preference  = "OnSuccess"
  force_update_tag    = "1"
  timeout             = "PT30M"

  script_content = <<EOF
          param([string] $name)
            $output = 'Hello {0}.' -f $name
            Write-Output $output
            $DeploymentScriptOutputs = @{}
            $DeploymentScriptOutputs['text'] = $output
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
