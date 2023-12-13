resource "aws_ce_anomaly_monitor" "pike" {
  name         = "AWSCustomAnomalyMonitor"
  monitor_type = "CUSTOM"

  monitor_specification = jsonencode({
    And            = null
    CostCategories = null
    Dimensions     = null
    Not            = null
    Or             = null

    Tags = {
      Key          = "CostCenter"
      MatchOptions = null
      Values = [
        "10000"
      ]
    }
  })

  tags = {
    pike = "permissions"
  }
}
