resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Resources/subscriptions/providers/read",
      //resources -azurerm_app_service_plan
      "Microsoft.Web/serverfarms/write",
      "Microsoft.Web/serverfarms/delete",
      "Microsoft.Web/serverfarms/read",

      //azurerm_app_service
      "Microsoft.Web/sites/write",
      "Microsoft.Web/sites/read",
      "Microsoft.Web/sites/delete",
      "Microsoft.Web/sites/config/write",
      "Microsoft.Web/sites/config/read",
      "Microsoft.Web/sites/config/delete",
      "Microsoft.Web/sites/config/list/action",
      "Microsoft.Web/sites/sourcecontrols/read",

      "Microsoft.Resources/subscriptions/providers/read",
      //azurerm_windows_web_app, azurerm_windows_function_app
      "Microsoft.Web/sites/read",
      //azurerm_app_service_plan
      "Microsoft.Web/serverfarms/read",
      //azurerm_public_ip_prefix
      "Microsoft.Network/publicIPPrefixes/read",
      //azurerm_public_ip
      "Microsoft.Network/publicIPAddresses/read",
      //azurerm_app_service_environment_v3, azurerm_app_service_environment
      "Microsoft.Web/hostingEnvironments/read",
      //azurerm_app_service_certificate_order
      "Microsoft.CertificateRegistration/certificateOrders/read",
      //azurerm_app_service_certificate
      "Microsoft.Web/certificates/read"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
