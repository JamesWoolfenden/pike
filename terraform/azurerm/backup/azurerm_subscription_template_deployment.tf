resource "azurerm_subscription_template_deployment" "pike_gen" {
  name             = "example-deployment"
  location         = "West Europe"
  template_content = <<TEMPLATE
 {
   "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
   "contentVersion": "1.0.0.0",
   "parameters": {},
   "variables": {},
   "resources": [
     {
       "type": "Microsoft.Resources/resourceGroups",
       "apiVersion": "2018-05-01",
       "location": "West Europe",
       "name": "some-resource-group",
       "properties": {}
     }
   ]
 }
 TEMPLATE

  // NOTE: whilst we show an inline template here, we recommend
  // sourcing this from a file for readability/editor support
}
