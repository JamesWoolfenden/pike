resource "azurerm_monitor_scheduled_query_rules_alert_v2" "pike_gen" {
  name                = "example-msqrv2"
  resource_group_name = "pike"
  location            = "pike"

  evaluation_frequency = "PT10M"
  window_duration      = "PT10M"
  scopes               = ["pike"]
  severity             = 4
  criteria {
    query                   = <<-QUERY
      requests
        | summarize CountByCountry=count() by client_CountryOrRegion
      QUERY
    time_aggregation_method = "Maximum"
    threshold               = 17.5
    operator                = "LessThan"

    resource_id_column    = "client_CountryOrRegion"
    metric_measure_column = "CountByCountry"
    dimension {
      name     = "client_CountryOrRegion"
      operator = "Exclude"
      values   = ["123"]
    }
    failing_periods {
      minimum_failing_periods_to_trigger_alert = 1
      number_of_evaluation_periods             = 1
    }
  }

  auto_mitigation_enabled          = true
  workspace_alerts_storage_enabled = false
  description                      = "example sqr"
  display_name                     = "example-sqr"
  enabled                          = true
  query_time_range_override        = "PT1H"
  skip_query_validation            = true
  action {
    action_groups = ["pike"]
    custom_properties = {
      key  = "value"
      key2 = "value2"
    }
    email_subject = "Email Header"
  }

  identity {
    type = "UserAssigned"
    identity_ids = [
      azurerm_user_assigned_identity.example.id,
    ]
  }
  tags = {
    key  = "value"
    key2 = "value2"
  }

  depends_on = ["pike"]
}
