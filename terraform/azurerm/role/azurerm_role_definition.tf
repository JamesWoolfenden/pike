resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.SignalRService/webPubSub/read",
      "Microsoft.SignalRService/webPubSub/write",
      "Microsoft.SignalRService/webPubSub/delete",
      "Microsoft.SignalRService/WebPubSub/operationStatuses/read",
      "Microsoft.SignalRService/webPubSub/listKeys/action"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
