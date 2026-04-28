resource "azurerm_application_insights_workbook_template" "pike_gen" {
  name                = "example-aiwt"
  resource_group_name = "pike"
  location            = "West Europe"
  author              = "test author"
  priority            = 1

  galleries {
    category      = "workbook"
    name          = "test"
    order         = 100
    resource_type = "microsoft.insights/components"
    type          = "tsg"
  }

  template_data = jsonencode({
    "version" : "Notebook/1.0",
    "items" : [
      {
        "type" : 1,
        "content" : {
          "json" : "## New workbook\n---\n\nWelcome to your new workbook."
        },
        "name" : "text - 2"
      }
    ],
    "styleSettings" : {},
    "$schema" : "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json"
  })

  localized = jsonencode({
    "ar" : [
      {
        "galleries" : [
          {
            "name" : "test",
            "category" : "Failures",
            "type" : "tsg",
            "resourceType" : "microsoft.insights/components",
            "order" : 100
          }
        ],
        "templateData" : {
          "version" : "Notebook/1.0",
          "items" : [
            {
              "type" : 1,
              "content" : {
                "json" : "## New workbook\n---\n\nWelcome to your new workbook."
              },
              "name" : "text - 2"
            }
          ],
          "styleSettings" : {},
          "$schema" : "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json"
        },
      }
    ]
  })

  tags = {
    key = "value"
  }
}
