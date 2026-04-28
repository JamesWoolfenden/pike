resource "azurerm_log_analytics_query_pack_query" "pike_gen" {
  name          = "19952bc3-0bf9-49eb-b713-6b80e7a41847"
  query_pack_id = "pike"
  body          = "let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"
  display_name  = "Exceptions - New in the last 24 hours"
}
