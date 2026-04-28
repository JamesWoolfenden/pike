resource "azurerm_logic_app_trigger_http_request" "pike_gen" {
  name         = "some-http-trigger"
  logic_app_id = "pike"

  schema = <<SCHEMA
{
    "type": "object",
    "properties": {
        "hello": {
            "type": "string"
        }
    }
}
SCHEMA

}
