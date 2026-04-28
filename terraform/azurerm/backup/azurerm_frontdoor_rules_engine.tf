resource "azurerm_frontdoor_rules_engine" "pike_gen" {
  name                = "exampleRulesEngineConfig1"
  frontdoor_name      = "pike"
  resource_group_name = "pike"

  rule {
    name     = "debuggingoutput"
    priority = 1

    action {
      response_header {
        header_action_type = "Append"
        header_name        = "X-TEST-HEADER"
        value              = "Append Header Rule"
      }
    }
  }

  rule {
    name     = "overwriteorigin"
    priority = 2

    match_condition {
      variable = "RequestMethod"
      operator = "Equal"
      value    = ["GET", "POST"]
    }

    action {

      response_header {
        header_action_type = "Overwrite"
        header_name        = "Access-Control-Allow-Origin"
        value              = "*"
      }

      response_header {
        header_action_type = "Overwrite"
        header_name        = "Access-Control-Allow-Credentials"
        value              = "true"
      }

    }
  }
}
