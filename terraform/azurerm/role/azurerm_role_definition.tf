resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.AppConfiguration/configurationStores/read",
      "Microsoft.AppConfiguration/configurationStores/write",
      "Microsoft.AppConfiguration/locations/deletedConfigurationStores/read",
      "Microsoft.AppConfiguration/configurationStores/listKeys/action",
      "Microsoft.AppConfiguration/configurationStores/delete"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
