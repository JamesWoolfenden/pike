resource "azurerm_resource_group_policy_assignment" "pike_gen" {
  name                 = "example"
  resource_group_id    = "pike"
  policy_definition_id = "pike"

  parameters = <<PARAMS
    {
      "tagName": {
        "value": "Business Unit"
      },
      "tagValue": {
        "value": "BU"
      }
    }
PARAMS
}
