resource "azurerm_log_analytics_solution" "pike" {
  solution_name         = "ContainerInsights"
  location              = "uksouth"
  resource_group_name   = "pike"
  workspace_resource_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.OperationalInsights/workspaces/pike"
  workspace_name        = "pike"

  plan {
    publisher = "Microsoft"
    product   = "OMSGallery/ContainerInsights"
  }
}
