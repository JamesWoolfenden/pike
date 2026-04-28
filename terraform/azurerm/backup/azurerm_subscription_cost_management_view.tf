resource "azurerm_subscription_cost_management_view" "pike_gen" {
  name         = "example"
  display_name = "Cost View per Month"
  chart_type   = "StackedColumn"
  accumulated  = false

  subscription_id = "/subscription/00000000-0000-0000-0000-000000000000"

  report_type = "Usage"
  timeframe   = "MonthToDate"

  dataset {
    granularity = "Monthly"

    aggregation {
      name        = "totalCost"
      column_name = "Cost"
    }
  }
  pivot {
    type = "Dimension"
    name = "ServiceName"
  }
  pivot {
    type = "Dimension"
    name = "ResourceLocation"
  }
  pivot {
    type = "Dimension"
    name = "ResourceGroupName"
  }
}
