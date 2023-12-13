resource "aws_ce_anomaly_subscription" "pike" {
  name      = "DAILYSUBSCRIPTION"
  frequency = "DAILY"

  monitor_arn_list = [
    aws_ce_anomaly_monitor.pike.arn,
  ]

  subscriber {
    type    = "EMAIL"
    address = "abc@example.com"
  }

  threshold_expression {
    dimension {
      key           = "ANOMALY_TOTAL_IMPACT_ABSOLUTE"
      values        = ["100.0"]
      match_options = ["GREATER_THAN_OR_EQUAL"]
    }
  }
  tags = {
    pike = "permissions"
  }

}
