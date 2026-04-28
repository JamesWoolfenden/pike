resource "azurerm_cosmosdb_sql_stored_procedure" "pike_gen" {
  name                = "test-stored-proc"
  resource_group_name = "pike"
  account_name        = "pike"
  database_name       = "pike"
  container_name      = "pike"

  body = <<BODY
   function () { var context = getContext(); var response = context.getResponse(); response.setBody('Hello, World'); }
BODY
}
