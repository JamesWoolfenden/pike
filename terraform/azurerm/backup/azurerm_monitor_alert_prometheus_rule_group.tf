resource "azurerm_monitor_alert_prometheus_rule_group" "pike_gen" {
  name                = "example-amprg"
  location            = "West Europe"
  resource_group_name = "pike"
  cluster_name        = "pike"
  description         = "This is the description of the following rule group"
  rule_group_enabled  = false
  interval            = "PT1M"
  scopes              = ["pike"]
  rule {
    enabled    = false
    expression = <<EOF
histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service="billing-processing"}[5m])) by (job_type))
EOF
    record     = "job_type:billing_jobs_duration_seconds:99p5m"
    labels = {
      team = "prod"
    }
  }

  rule {
    alert      = "Billing_Processing_Very_Slow"
    enabled    = true
    expression = <<EOF
histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service="billing-processing"}[5m])) by (job_type))
EOF
    for        = "PT5M"
    severity   = 2

    action {
      action_group_id = "pike"
    }

    alert_resolution {
      auto_resolved   = true
      time_to_resolve = "PT10M"
    }

    annotations = {
      annotationName = "annotationValue"
    }

    labels = {
      team = "prod"
    }
  }
  tags = {
    key = "value"
  }
}
