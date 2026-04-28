resource "azurerm_cost_management_scheduled_action" "pike_gen" {
  name         = "examplescheduledaction"
  display_name = "Report Last 6 Months"

  view_id = "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/views/ms:CostByService"

  email_address_sender = "platformteam@test.com"
  email_subject        = "Cost Management Report"
  email_addresses      = ["example@example.com"]
  message              = "Hi all, take a look at last 6 months spending!"

  frequency  = "Daily"
  start_date = "2023-01-02T00:00:00Z"
  end_date   = "2023-02-02T00:00:00Z"
}
