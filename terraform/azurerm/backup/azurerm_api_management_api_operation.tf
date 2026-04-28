resource "azurerm_api_management_api_operation" "pike_gen" {
  operation_id        = "user-delete"
  api_name            = "pike"
  api_management_name = "pike"
  resource_group_name = "pike"
  display_name        = "Delete User Operation"
  method              = "DELETE"
  url_template        = "/users/{id}/delete"
  description         = "This can only be done by the logged in user."

  template_parameter {
    name     = "id"
    type     = "number"
    required = true
  }

  response {
    status_code = 200
  }
}
